package repos

import (
	"context"
	"fmt"
	"strings"

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
		autoCategoryIDs []int64,
		filters types.BiddingFilters) (*types.BiddingAutoResultSet, error)
	CreateBidding(
		ctx context.Context,
		userID, companyID int64,
		bidding types.Bidding,
		biddingItems []struct{ types.BiddingItem }) error
	GetBidding(
		ctx context.Context,
		biddingID int64,
	) (*types.BiddingBiddingItems, error)
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
	err := r.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM biddings WHERE company_id = $1", companyID)
	if err != nil {
		log.Error().Err(err).Msg("Error counting biddings")
		return nil, err
	}
	result.Count = count

	var biddings []types.Bidding
	err = r.db.SelectContext(ctx, &biddings, `
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

	for _, bidding := range biddings {
		var biddingItems []types.BiddingItem
		err := r.db.SelectContext(ctx, &biddingItems, `
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
		result.Data = append(result.Data, types.BiddingBiddingItems{
			Bidding: bidding,
			Items:   biddingItems,
		})
	}

	return &result, nil
}

func (r *PgBiddingRepository) GetAllBiddingsByAutoCategoryIDs(
	ctx context.Context,
	autoCategoryIDs []int64,
	filters types.BiddingFilters,
) (*types.BiddingAutoResultSet, error) {
	result := types.BiddingAutoResultSet{}
	acIN := []string{}
	for _, acID := range autoCategoryIDs {
		acIN = append(acIN, fmt.Sprintf("%d", acID))
	}
	var count int
	err := r.db.GetContext(ctx, &count, `
  SELECT COUNT(DISTINCT b.id) 
  FROM biddings b 
  LEFT JOIN bidding_items bi ON b.id = bi.bidding_id
  WHERE bi.auto_category_id IN (`+strings.Join(acIN, ",")+`)
  AND b.status NOT IN ('CANCELED', 'FINISHED');
  `)
	if err != nil {
		log.Error().Err(err).Msg("Error counting biddings for auto companies")
		return nil, err
	}
	result.Count = count

	var biddings []types.Bidding
	err = r.db.SelectContext(ctx, &biddings, `
		SELECT b.id, b.customer_name, b.vehicle_brand, b.vehicle_name, b.vehicle_year, b.vehicle_color, 
    COALESCE(b.notes, '') as notes, b.status, c.telephone as company_phone, b.created_at, b.updated_at
		FROM biddings b
    LEFT JOIN companies c ON c.id = b.company_id
    LEFT JOIN bidding_items bi ON b.id = bi.bidding_id
		WHERE bi.auto_category_id IN (`+strings.Join(acIN, ",")+`)
    AND b.status NOT IN ('CANCELED', 'FINISHED')
    GROUP BY b.id, c.telephone
    ORDER BY `+filters.OrderBy+` `+filters.Order+`
    LIMIT $1 OFFSET $2`,
		filters.Limit, filters.Offset)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting biddings")
		return nil, err
	}

	var biddingItems []types.BiddingItem
	for _, bidding := range biddings {
		err := r.db.SelectContext(ctx, &biddingItems, `
    SELECT status, created_at, updated_at, COALESCE(notes, '') as notes,
    ac.description as auto_category_description, ac.type as auto_category_type
    FROM bidding_items bi
    LEFT JOIN auto_categories ac ON ac.id = bi.auto_category_id
    WHERE bi.bidding_id = $1
    AND bi.auto_category_id IN (`+strings.Join(acIN, ",")+`)
    AND bi.status NOT IN ('CANCELED')
    `, bidding.ID)
		if err != nil {
			log.Error().Err(err).Msg("Error selecting bidding_items for a bidding")
			return nil, err
		}

		var offers []types.Offer
		var biddingItemOffers []types.BiddingItemOffers
		for _, biddingItem := range biddingItems {
			err := r.db.SelectContext(ctx, &offers, `
    SELECT price, created_at, updated_at, COALESCE(notes, '') as notes
    FROM offers
    WHERE bidding_item_id = $1;
    `, biddingItem.ID)
			if err != nil {
				log.Error().Err(err).Msg("Error selecting offers for a bidding_item")
				return nil, err
			}
			biddingItemOffers = append(biddingItemOffers, types.BiddingItemOffers{
				BiddingItem: biddingItem,
				Offers:      offers,
			})
		}
		result.Data = append(result.Data, types.BiddingBiddingItemsOffers{
			Bidding: bidding,
			Items:   biddingItemOffers,
		})
	}

	return &result, nil
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
    RETURNING *;
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

func (r *PgBiddingRepository) GetBidding(ctx context.Context, biddingID int64) (*types.BiddingBiddingItems, error) {
	var biddings types.Bidding
	err := r.db.GetContext(ctx, &biddings, `
		SELECT id, customer_name, vehicle_brand, vehicle_name, vehicle_year, vehicle_color, 
    COALESCE(notes, '') as notes, status, created_at, updated_at
		FROM biddings
		WHERE id = $1;
    `, biddingID)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting biddings")
		return nil, err
	}

	var biddingItems []types.BiddingItem
	err = r.db.SelectContext(ctx, &biddingItems, `
    SELECT bi.id, bi.status, bi.created_at, bi.updated_at, COALESCE(bi.notes, '') as notes, 
    ac.description as auto_category_description, ac.type as auto_category_type
    FROM bidding_items bi
    LEFT JOIN auto_categories ac ON bi.auto_category_id = ac.id
    WHERE bi.bidding_id = $1;
    `, biddings.ID)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting bidding_items")
		return nil, err
	}

	return &types.BiddingBiddingItems{
		Bidding: biddings,
		Items:   biddingItems,
	}, nil
}
