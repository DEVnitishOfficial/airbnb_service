package db

import (
	"context"
	"database/sql"
	"fmt"
)

type reviewRepositoryImpl struct {
	db *sql.DB
}

// NewReviewRepository creates a new repository instance.
func NewReviewRepositoryAggRating(db *sql.DB) ReviewRepositoryAggRating {
	return &reviewRepositoryImpl{db: db}
}

// FetchUnappliedAggregates groups unapplied ratings by hotel.
func (r *reviewRepositoryImpl) FetchUnappliedAggregates(ctx context.Context, cutoff string) ([]AggregatedRating, error) {
	query := `
		SELECT hotel_id, SUM(rating) AS total_rating, COUNT(*) AS cnt
		FROM reviews
		WHERE is_synced = FALSE AND created_at <= ?
		GROUP BY hotel_id
	`

	rows, err := r.db.QueryContext(ctx, query, cutoff)
	if err != nil {
		return nil, fmt.Errorf("query aggregates: %w", err)
	}
	defer rows.Close()

	var results []AggregatedRating
	for rows.Next() {
		var a AggregatedRating
		var sum sql.NullFloat64
		var cnt sql.NullInt64

		if err := rows.Scan(&a.HotelID, &sum, &cnt); err != nil {
			return nil, fmt.Errorf("scan agg row: %w", err)
		}
		if sum.Valid {
			a.Sum = sum.Float64
		}
		if cnt.Valid {
			a.Count = cnt.Int64
		}
		results = append(results, a)
	}
	return results, rows.Err()
}

// MarkReviewsAsSynced updates processed reviews.
func (r *reviewRepositoryImpl) MarkReviewsAsSynced(ctx context.Context, tx *sql.Tx, hotelID int64, cutoff string) error {
	_, err := tx.ExecContext(ctx, `
		UPDATE reviews 
		SET is_synced = TRUE 
		WHERE is_synced = FALSE AND hotel_id = ? AND created_at <= ?`, hotelID, cutoff)
	return err
}

// BeginTx starts a transaction.
func (r *reviewRepositoryImpl) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}
