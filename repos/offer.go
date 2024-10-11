package repos

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/gmtborges/orcamento-auto/types"
)

type OfferRepository interface {
	GetOffersByBiddingItemID(ctx context.Context, biddingItemID int64) ([]types.Offer, error)
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
) ([]types.Offer, error) {
	var offers []types.Offer
	err := r.db.SelectContext(ctx, &offers, `SELECT o.*, c.name as company_name
  FROM offers o 
  LEFT JOIN companies c ON o.company_id = c.id
  WHERE o.bidding_item_id = $1
  `, biddingItemID)
	if err != nil {
		return nil, err
	}

	return offers, nil
}
