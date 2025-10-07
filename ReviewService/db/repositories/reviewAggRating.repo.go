package db

import (
	"context"
	"database/sql"
)

// AggregatedRating holds grouped rating data for a hotel.
type AggregatedRating struct {
	HotelID int64
	Sum     float64
	Count   int64
}

// ReviewRepositoryAggRating defines all DB methods related to reviews.
type ReviewRepositoryAggRating interface {
	FetchUnappliedAggregates(ctx context.Context, cutoff string) ([]AggregatedRating, error)
	MarkReviewsAsSynced(ctx context.Context, tx *sql.Tx, hotelID int64, cutoff string) error
	BeginTx(ctx context.Context) (*sql.Tx, error)
}
