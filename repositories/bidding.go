package repositories

import (
	"context"
	"database/sql"

	"github.com/gmtborges/orcamento-auto/models"
)

type BiddingRepository interface {
	AllBiddings(ctx context.Context, userID int) (models.BiddingSlice, error)
}

type PostgresBiddingRepository struct {
	db *sql.DB
}

func NewPostgresBiddingRepository(db *sql.DB) *PostgresBiddingRepository {
	return &PostgresBiddingRepository{db: db}
}

func (r *PostgresBiddingRepository) AllBiddings(ctx context.Context, userID int) (models.BiddingSlice, error) {
	biddings, err := models.Biddings().All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return biddings, nil
}
