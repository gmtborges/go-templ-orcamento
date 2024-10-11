package repos

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/types"
)

type AutoCategoryRepository interface {
	GetAllAutoCategories(ctx context.Context) ([]types.AutoCategory, error)
	GetAllAutoCategoryIDsByCompanyID(ctx context.Context, companyID int64) ([]int64, error)
}

type PgAutoCategoryRepository struct {
	db *sqlx.DB
}

func NewPgAutoCategoryRepository(db *sqlx.DB) *PgAutoCategoryRepository {
	return &PgAutoCategoryRepository{db: db}
}

func (r *PgAutoCategoryRepository) GetAllAutoCategories(ctx context.Context) ([]types.AutoCategory, error) {
	var ac []types.AutoCategory
	err := r.db.SelectContext(ctx, &ac, "SELECT * FROM auto_categories")
	if err != nil {
		log.Error().Err(err).Msg("Error selecting auto_categories.")
	}
	return ac, err
}

func (r *PgAutoCategoryRepository) GetAllAutoCategoryIDsByCompanyID(ctx context.Context, companyID int64) ([]int64, error) {
	var acIDs []int64
	err := r.db.SelectContext(ctx, &acIDs, `SELECT auto_category_id 
  FROM companies_auto_categories
  WHERE company_id = $1;
  `, companyID)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting companies_auto_categories.")
	}
	return acIDs, err
}
