package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/types"
)

type BiddingRepository interface {
	AllBiddings(ctx context.Context, companyID int64, filters types.BiddingFilters) (*types.BiddingResult, error)
	CreateBidding(ctx context.Context, companyID int64, bidding *types.Bidding, biddingItems []types.BiddingItem) (*types.Bidding, error)
}

type PostgresBiddingRepository struct {
	db *sqlx.DB
}

func NewPostgresBiddingRepository(db *sqlx.DB) *PostgresBiddingRepository {
	return &PostgresBiddingRepository{db: db}
}

func (r *PostgresBiddingRepository) AllBiddings(ctx context.Context, companyID int64, filters types.BiddingFilters) (*types.BiddingResult, error) {
	result := types.BiddingResult{}

	count := 0
	err := r.db.Get(&count, "SELECT COUNT(*) FROM biddings")
	if err != nil {
		log.Error().Err(err).Msg("Error counting items")
		return nil, err
	}
	result.Count = count

	biddings := []types.Bidding{}
	err = r.db.Select(&biddings, `
		SELECT *
		FROM biddings
		WHERE company_id = $1
    ORDER BY `+filters.OrderBy+` `+filters.Order+`
    LIMIT $2 OFFSET $3`,
		companyID, filters.Limit, filters.Offset)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting biddings")
		return nil, err
	}

	for _, bidding := range biddings {
		biddingItems := []types.BiddingItemModel{}
		err = r.db.Select(&biddingItems, `
    SELECT bi.*, ac.id as auto_category_id, ac.name as auto_category_name, ac.type as auto_category_type
    FROM bidding_items bi
    LEFT JOIN auto_categories ac ON bi.auto_category_id = ac.id
    WHERE bidding_id = $1
    `, bidding.ID)
		if err != nil {
			log.Error().Err(err).Msg("Error selecting bidding items")
			return nil, err
		}
		result.Data = append(result.Data, types.BiddingModel{
			Bidding: bidding,
			Items:   biddingItems,
		})
	}

	return &result, nil
}

func (r *PostgresBiddingRepository) CreateBidding(ctx context.Context, companyID int64, bidding *types.Bidding, biddingItems []types.BiddingItem) (*types.Bidding, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error creating transaction")
		return nil, err
	}
	defer tx.Rollback()

	err = tx.Get(bidding, `
    INSERT INTO biddings (user_id, company_id, customer_name, description, created_at, updated_at)
    VALUES ($1, $2, $3, now(), now())
    RETURNING *
  `, companyID, bidding.CustomerName, bidding.Description)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting bidding")
		return nil, err
	}

	for _, item := range biddingItems {
		_, err := tx.Exec(`
      INSERT INTO bidding_items (bidding_id, auto_category_id)
      VALUES ($1, $2, $3, $4)
    `, bidding.ID, item.AutoCategoryID)
		if err != nil {
			log.Error().Err(err).Msg("Error inserting bidding item")
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msg("Error committing transaction")
		return nil, err
	}

	return bidding, nil
}
