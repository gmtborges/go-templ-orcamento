package repo

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/gmtborges/orcamento-auto/types"
)

type PropostaRepository interface {
	GetPropostasByOrcamentoItemID(ctx context.Context, oiID int64) ([]types.PropostaModel, error)
}

type PgPropostaRepository struct {
	db *sqlx.DB
}

func NewPgPropostaRepository(db *sqlx.DB) *PgPropostaRepository {
	return &PgPropostaRepository{db: db}
}

func (r *PgPropostaRepository) GetPropostasByOrcamentoItemID(
	ctx context.Context,
	oiID int64,
) ([]types.PropostaModel, error) {
	var p []types.PropostaModel
	err := r.db.SelectContext(ctx, &p, `SELECT p.*, e.nome as empresa_nome 
  FROM propostas p 
  LEFT JOIN empresas e ON p.empresa_id = e.id
  LEFT JOIN orcamento_itens oi oi ON oi.id = p.orcamento_item_id
  WHERE oi.id = $1
  `, oiID)
	if err != nil {
		return nil, err
	}

	return p, nil
}
