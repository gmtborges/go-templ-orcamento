package repositories

import (
	"context"
	"database/sql"

	"github.com/gustavomtborges/orcamento-auto/models"
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

func (st *PostgresBiddingRepository) AllBiddings(ctx context.Context, userID int) (models.BiddingSlice, error) {
	biddings, err := models.Biddings().All(ctx, st.db)
	if err != nil {
		return nil, err
	}
	return biddings, nil
}
