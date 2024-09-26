package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/types"
)

type AutoCategoriaRepository interface {
	GetAllAutoCategorias(ctx context.Context) ([]types.AutoCategoria, error)
}

type PgAutoCategoriaRepository struct {
	db *sqlx.DB
}

func NewPgAutoCategoriaRepository(db *sqlx.DB) *PgAutoCategoriaRepository {
	return &PgAutoCategoriaRepository{db: db}
}

func (r *PgAutoCategoriaRepository) GetAllAutoCategorias(ctx context.Context) ([]types.AutoCategoria, error) {
	var ac []types.AutoCategoria
	err := r.db.SelectContext(ctx, &ac, "SELECT * FROM auto_categorias")
	if err != nil {
		log.Error().Err(err).Msg("Error selecting auto_categorias.")
	}
	return ac, err
}
