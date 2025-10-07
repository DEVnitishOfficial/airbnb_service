package services

import (
	clients "ReviewService/client"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type ReviewBatchProcessor interface {
	ProcessPendingRatings(ctx context.Context) error
}

// ReviewBatchProcessorImpl contains DB handle and configuration for batch processing
type ReviewBatchProcessorImpl struct {
	DB          *sql.DB
	LockName    string // name for MySQL GET_LOCK
	HotelClient *clients.HotelClient
}

// NewReviewBatchProcessor creates a new instance
func NewReviewBatchProcessor(db *sql.DB, hotelClient *clients.HotelClient) *ReviewBatchProcessorImpl {
	return &ReviewBatchProcessorImpl{
		DB:          db,
		LockName:    "process_pending_ratings",
		HotelClient: hotelClient,
	}
}

// s is the instance of ReviewBatchProcessorImpl, and we use it to access the struct's fields (like DB and HotelClient) to perform operations needed for processing ratings.
func (s *ReviewBatchProcessorImpl) ProcessPendingRatings(ctx context.Context) error {
	var got sql.NullInt64
	// Try to get a named lock called process_pending_ratings from MySQL.
	// If we successfully get it within 1 second, continue. If not, skip execution because someone else is already running the batch.
	if err := s.DB.QueryRowContext(ctx, "SELECT GET_LOCK(?, 1)", s.LockName).Scan(&got); err != nil {
		return fmt.Errorf("GET_LOCK error: %w", err)
	}
	if !got.Valid || got.Int64 != 1 {
		// Could not acquire lock (someone else is running). Exit gracefully.
		log.Println("ProcessPendingRatings: another instance is already running")
		return nil
	}
	// This block ensures that no matter how the function exits (success, error, panic, or early return), the MySQL advisory lock is released.
	defer func() {
		if _, err := s.DB.ExecContext(context.Background(), "SELECT RELEASE_LOCK(?)", s.LockName); err != nil {
			log.Println("error releasing lock:", err)
		}
	}()

	// Cutoff: process reviews created up to the current hour boundary.
	// Truncate just rounds off to the hour, so if now is 10:34, cutoff is 10:00.
	cutoff := time.Now().UTC()
	log.Printf("ProcessPendingRatings: running cutoff=%s\n", cutoff.Format(time.RFC3339))

	// Step A: fetch aggregates grouped by hotel_id
	rows, err := s.DB.QueryContext(ctx, `
SELECT hotel_id, SUM(rating) AS total_rating, COUNT(*) AS cnt
FROM reviews
WHERE is_synced = FALSE AND created_at <= ?
GROUP BY hotel_id
`, cutoff)
	if err != nil {
		return fmt.Errorf("query aggregates: %w", err)
	}
	defer rows.Close()

	type agg struct {
		HotelID int64
		Sum     float64
		Count   int64
	}

	var aggs []agg
	for rows.Next() {
		var a agg
		var sum sql.NullFloat64
		var cnt sql.NullInt64
		if err := rows.Scan(&a.HotelID, &sum, &cnt); err != nil {
			return fmt.Errorf("scan agg row: %w", err)
		}
		if sum.Valid {
			a.Sum = sum.Float64
		}
		if cnt.Valid {
			a.Count = cnt.Int64
		}
		aggs = append(aggs, a)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("rows err: %w", err)
	}

	if len(aggs) == 0 {
		log.Println("ProcessPendingRatings: no unapplied reviews found")
		return nil
	}

	// Process each hotel in its own transaction to limit lock scope and potential rollback size.
	for _, a := range aggs {
		// start tx
		tx, err := s.DB.BeginTx(ctx, nil)
		if err != nil {
			log.Printf("begin tx hotel %d: %v\n", a.HotelID, err)
			continue
		}

		hotelData, err := s.HotelClient.GetHotelRating(a.HotelID)
		if err != nil {
			tx.Rollback()
			log.Printf("hotel fetch failed for %d: %v\n", a.HotelID, err)
			continue
		}

		oldAvg, _ := hotelData.Rating.Float64()
		oldCnt := hotelData.RatingCount

		newCount := oldCnt + a.Count
		var newAvg float64
		if newCount > 0 {
			// Compute new average rating
			// newAvg = (oldAvg*oldCnt + sum_of_new_ratings) / newCount
			// where sum_of_new_ratings is the sum of ratings from the new reviews being applied
			// This formula ensures that we correctly weight the old average by its count and add the new ratings.
			// newAvg must be in in the range [1.0, 5.0], and also with correct decimal precision.

			newAvg = ((oldAvg * float64(oldCnt)) + a.Sum) / float64(newCount)

		} else {
			newAvg = 0
		}

		// Update hotels table
		if err := s.HotelClient.UpdateHotelRating(a.HotelID, newAvg, newCount); err != nil {
			tx.Rollback()
			log.Printf("hotel update failed for %d: %v\n", a.HotelID, err)
			continue
		}

		// Mark the used reviews as synced
		if _, err := tx.ExecContext(ctx, "UPDATE reviews SET is_synced = TRUE WHERE is_synced = FALSE AND hotel_id = ? AND created_at <= ?", a.HotelID, cutoff); err != nil {
			tx.Rollback()
			log.Printf("mark reviews synced for hotel %d err: %v\n", a.HotelID, err)
			continue
		}

		if err := tx.Commit(); err != nil {
			log.Printf("commit err for hotel %d: %v\n", a.HotelID, err)
			// best-effort: continue to next
			continue
		}

		log.Printf("Processed hotel %d: added %d ratings (sum=%.2f) -> new_count=%d new_avg=%.2f\n", a.HotelID, a.Count, a.Sum, newCount, newAvg)
	}

	return nil

}
