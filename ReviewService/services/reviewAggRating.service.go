package services

import (
	client "ReviewService/client"
	repositories "ReviewService/db/repositories"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type ReviewBatchProcessor interface {
	ProcessPendingRatings(ctx context.Context) error
}

type ReviewBatchProcessorImpl struct {
	repo        repositories.ReviewRepositoryAggRating
	hotelClient *client.HotelClient
	db          *sql.DB
	lockName    string
}

func NewReviewBatchProcessor(db *sql.DB, repo repositories.ReviewRepositoryAggRating, hotelClient *client.HotelClient) *ReviewBatchProcessorImpl {
	return &ReviewBatchProcessorImpl{
		repo:        repo,
		hotelClient: hotelClient,
		db:          db,
		lockName:    "process_pending_ratings",
	}
}

func (s *ReviewBatchProcessorImpl) ProcessPendingRatings(ctx context.Context) error {
	// acquire MySQL lock
	var got sql.NullInt64
	if err := s.db.QueryRowContext(ctx, "SELECT GET_LOCK(?, 1)", s.lockName).Scan(&got); err != nil {
		return fmt.Errorf("GET_LOCK error: %w", err)
	}
	if !got.Valid || got.Int64 != 1 {
		log.Println("ProcessPendingRatings: another instance is already running")
		return nil
	}
	defer s.db.ExecContext(context.Background(), "SELECT RELEASE_LOCK(?)", s.lockName)

	cutoff := time.Now().UTC().Format("2006-01-02 15:04:05")
	log.Printf("🕒 Running batch with cutoff=%s\n", cutoff)

	aggs, err := s.repo.FetchUnappliedAggregates(ctx, cutoff)
	if err != nil {
		return fmt.Errorf("fetch unapplied aggregates: %w", err)
	}
	if len(aggs) == 0 {
		log.Println("No unapplied reviews found")
		return nil
	}

	for _, a := range aggs {
		tx, err := s.repo.BeginTx(ctx)
		if err != nil {
			log.Printf("begin tx hotel %d: %v", a.HotelID, err)
			continue
		}

		hotelData, err := s.hotelClient.GetHotelRating(a.HotelID)
		if err != nil {
			tx.Rollback()
			log.Printf("hotel fetch failed for %d: %v", a.HotelID, err)
			continue
		}

		oldAvg, _ := hotelData.Rating.Float64()
		oldCnt := hotelData.RatingCount
		newCount := oldCnt + a.Count
		// weighted average calculation
		// newAvg = (oldAvg * oldCnt + a.Sum) / newCount
		// where a.Sum is the sum of new ratings, a.Count is the count of new ratings
		// and newCount = oldCnt + a.Count
		// This ensures that larger counts have more influence on the new average.
		newAvg := ((oldAvg * float64(oldCnt)) + a.Sum) / float64(newCount)

		if err := s.hotelClient.UpdateHotelRating(a.HotelID, newAvg, newCount); err != nil {
			tx.Rollback()
			log.Printf("hotel update failed for %d: %v", a.HotelID, err)
			continue
		}

		if err := s.repo.MarkReviewsAsSynced(ctx, tx, a.HotelID, cutoff); err != nil {
			tx.Rollback()
			log.Printf("mark reviews synced for hotel %d err: %v", a.HotelID, err)
			continue
		}

		if err := tx.Commit(); err != nil {
			log.Printf("commit err for hotel %d: %v", a.HotelID, err)
			continue
		}

		log.Printf("✅ Processed hotel %d: added %d ratings (sum=%.2f) → new_count=%d new_avg=%.2f",
			a.HotelID, a.Count, a.Sum, newCount, newAvg)
	}
	return nil
}
