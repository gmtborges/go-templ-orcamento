package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/gmtborges/orcamento-auto/types"
)

type BiddingRepository interface {
	AllBiddings(ctx context.Context, userID int) (*types.BiddingSlice, error)
}

type PostgresBiddingRepository struct {
	db *sqlx.DB
}

func NewPostgresBiddingRepository(db *sqlx.DB) *PostgresBiddingRepository {
	return &PostgresBiddingRepository{db: db}
}

func (r *PostgresBiddingRepository) AllBiddings(ctx context.Context, userID int) (*types.BiddingSlice, error) {
	biddingSlide := types.BiddingSlice{}
	err := r.db.Select(&biddingSlide, "SELECT * FROM biddings WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return &biddingSlide, nil
}
