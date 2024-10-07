package repos

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/types"
)

type BiddingRepository interface {
	GetAllBiddingsByCompanyID(
		ctx context.Context,
		companyID int64,
		filters types.BiddingFilters) (*types.BiddingResultSet, error)
	GetAllBiddingsByAutoCategoryIDs(
		ctx context.Context,
		autoCategoryIDs []int64) (*types.BiddingAutoResultSet, error)
	CreateBidding(
		ctx context.Context,
		userID, companyID int64,
		bidding types.Bidding,
		biddingItems []struct{ types.BiddingItem }) error
	GetBidding(
		ctx context.Context,
		biddingID int64,
	) (*types.BiddingModel, error)
}

type PgBiddingRepository struct {
	db *sqlx.DB
}

func NewPgBiddingRepository(db *sqlx.DB) *PgBiddingRepository {
	return &PgBiddingRepository{db: db}
}

func (r *PgBiddingRepository) GetAllBiddingsByCompanyID(
	ctx context.Context,
	companyID int64,
	filters types.BiddingFilters,
) (*types.BiddingResultSet, error) {
	result := types.BiddingResultSet{}

	var count int
	err := r.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM biddings")
	if err != nil {
		log.Error().Err(err).Msg("Error counting biddings")
		return nil, err
	}
	result.Count = count

	b := []types.Bidding{}
	err = r.db.SelectContext(ctx, &b, `
		SELECT id, customer_name, vehicle_brand, vehicle_name, vehicle_year, vehicle_color, 
    COALESCE(notes, '') as notes, status, created_at, updated_at
		FROM biddings
		WHERE company_id = $1
    ORDER BY `+filters.OrderBy+` `+filters.Order+`
    LIMIT $2 OFFSET $3`,
		companyID, filters.Limit, filters.Offset)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting biddings")
		return nil, err
	}

	for _, bidding := range b {
		var bi []types.BiddingItemModel
		err := r.db.SelectContext(ctx, &bi, `
    SELECT bi.status, bi.created_at, bi.updated_at, COALESCE(bi.notes, '') as notes, 
    ac.description as auto_category_description, ac.type as auto_category_type
    FROM bidding_items bi
    LEFT JOIN auto_categories ac ON bi.auto_category_id = ac.id
    WHERE bi.bidding_id = $1
    `, bidding.ID)
		if err != nil {
			log.Error().Err(err).Msg("Error selecting bidding_items")
			return nil, err
		}
		result.Data = append(result.Data, types.BiddingModel{
			Bidding: bidding,
			Items:   bi,
		})
	}

	return &result, nil
}

func (r *PgBiddingRepository) GetAllBiddingsByAutoCategoryIDs(
	ctx context.Context,
	autoCategoryIDs []int64,
) (*types.BiddingAutoResultSet, error) {
	return nil, nil
}

func (r *PgBiddingRepository) CreateBidding(
	ctx context.Context,
	userID, companyID int64,
	bidding types.Bidding,
	biddingItems []struct{ types.BiddingItem },
) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error creating transaction")
		return err
	}
	defer tx.Rollback()

	err = tx.Get(&bidding, `
    INSERT INTO biddings 
    (user_id, company_id, customer_name, vehicle_brand, vehicle_name, 
    vehicle_year, vehicle_color, notes, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now(), now())
    RETURNING *
  `, userID, companyID, bidding.CustomerName, bidding.VehicleBrand,
		bidding.VehicleName, bidding.VehicleYear, bidding.VehicleColor, bidding.Notes)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting into biddings")
		return err
	}

	for _, item := range biddingItems {
		_, err := tx.Exec(`
      INSERT INTO bidding_items (bidding_id, auto_category_id, notes, created_at, updated_at)
      VALUES ($1, $2, $3, now(), now())
    `, bidding.ID, item.AutoCategoryID, item.Notes)
		if err != nil {
			log.Error().Err(err).Msg("Error inserting into bidding_items")
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msg("Error committing transaction")
		return err
	}

	return nil
}

func (r *PgBiddingRepository) GetBidding(ctx context.Context, biddingID int64) (*types.BiddingModel, error) {
	b := types.Bidding{}
	err := r.db.GetContext(ctx, &b, `
		SELECT id, customer_name, vehicle_brand, vehicle_name, vehicle_year, vehicle_color, 
    COALESCE(notes, '') as notes, status, created_at, updated_at
		FROM biddings
		WHERE id = $1`, biddingID)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting biddings")
		return nil, err
	}

	var bi []types.BiddingItemModel
	err = r.db.SelectContext(ctx, &bi, `
    SELECT bi.id, bi.status, bi.created_at, bi.updated_at, COALESCE(bi.notes, '') as notes, 
    ac.description as auto_category_description, ac.type as auto_category_type
    FROM bidding_items bi
    LEFT JOIN auto_categories ac ON bi.auto_category_id = ac.id
    WHERE bi.bidding_id = $1
    `, b.ID)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting bidding_items")
		return nil, err
	}

	return &types.BiddingModel{
		Bidding: b,
		Items:   bi,
	}, nil
}
