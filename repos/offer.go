package repos

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/gmtborges/orcamento-auto/types"
)

type OfferRepository interface {
	GetOffersByBiddingItemID(ctx context.Context, biddingItemID int64) ([]types.OfferModel, error)
}

type PgOfferRepository struct {
	db *sqlx.DB
}

func NewPgOfferRepository(db *sqlx.DB) *PgOfferRepository {
	return &PgOfferRepository{db: db}
}

func (r *PgOfferRepository) GetOffersByBiddingItemID(
	ctx context.Context,
	biddingItemID int64,
) ([]types.OfferModel, error) {
	var p []types.OfferModel
	err := r.db.SelectContext(ctx, &p, `SELECT o.*, c.name as company_name
  FROM offers o 
  LEFT JOIN companies c ON o.company_id = c.id
  LEFT JOIN bidding_items bi bi ON bi.id = o.bidding_item_id
  WHERE bi.id = $1
  `, biddingItemID)
	if err != nil {
		return nil, err
	}

	return p, nil
}
