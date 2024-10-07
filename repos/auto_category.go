package repos

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/types"
)

type AutoCategoryRepository interface {
	GetAllAutoCategories(ctx context.Context) ([]types.AutoCategory, error)
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
