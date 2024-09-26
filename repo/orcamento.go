package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/types"
)

type OrcamentoRepository interface {
	GetAllOrcamentos(
		ctx context.Context,
		empID int64,
		filtros types.OrcamentoFiltros) (*types.OrcamentoResultSet, error)
	CreateOrcamento(
		ctx context.Context,
		uID, empID int64,
		orcamento types.Orcamento,
		orcamentoItens []struct{ types.OrcamentoItem }) error
	GetOrcamento(
		ctx context.Context,
		orcamentoID int64,
	) (*types.OrcamentoModel, error)
}

type PgOrcamentoRepository struct {
	db *sqlx.DB
}

func NewPgOrcamentoRepository(db *sqlx.DB) *PgOrcamentoRepository {
	return &PgOrcamentoRepository{db: db}
}

func (r *PgOrcamentoRepository) GetAllOrcamentos(
	ctx context.Context,
	empID int64,
	filtros types.OrcamentoFiltros,
) (*types.OrcamentoResultSet, error) {
	result := types.OrcamentoResultSet{}

	var count int
	err := r.db.GetContext(ctx, &count, "SELECT COUNT(*) FROM orcamentos")
	if err != nil {
		log.Error().Err(err).Msg("Error counting orcamentos")
		return nil, err
	}
	result.Count = count

	o := []types.Orcamento{}
	err = r.db.SelectContext(ctx, &o, `
		SELECT id, associado_nome, veiculo_marca, veiculo_nome, veiculo_ano, veiculo_cor, 
    COALESCE(observacao, '') as observacao, status, data_criacao, data_atualizacao
		FROM orcamentos
		WHERE empresa_id = $1
    ORDER BY `+filtros.OrderBy+` `+filtros.Order+`
    LIMIT $2 OFFSET $3`,
		empID, filtros.Limit, filtros.Offset)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting orcamentos")
		return nil, err
	}

	for _, orcamento := range o {
		var oi []types.OrcamentoItemModel
		err := r.db.SelectContext(ctx, &oi, `
    SELECT oi.status, oi.data_criacao, oi.data_atualizacao, COALESCE(oi.observacao, '') as observacao, 
    ac.descricao as auto_categoria_descricao, ac.tipo as auto_categoria_tipo
    FROM orcamento_itens oi
    LEFT JOIN auto_categorias ac ON oi.auto_categoria_id = ac.id
    WHERE oi.orcamento_id = $1
    `, orcamento.ID)
		if err != nil {
			log.Error().Err(err).Msg("Error selecting orcamento_itens")
			return nil, err
		}
		result.Data = append(result.Data, types.OrcamentoModel{
			Orcamento: orcamento,
			Itens:     oi,
		})
	}

	return &result, nil
}

func (r *PgOrcamentoRepository) CreateOrcamento(
	ctx context.Context,
	uID, empID int64,
	orcamento types.Orcamento,
	orcamentoItems []struct{ types.OrcamentoItem },
) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("Error creating transaction")
		return err
	}
	defer tx.Rollback()

	err = tx.Get(&orcamento, `
    INSERT INTO orcamentos 
    (usuario_id, empresa_id, associado_nome, veiculo_marca, veiculo_nome, 
    veiculo_ano, veiculo_cor, observacao, data_criacao, data_atualizacao)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now(), now())
    RETURNING *
  `, uID, empID, orcamento.AssociadoNome, orcamento.VeiculoMarca,
		orcamento.VeiculoNome, orcamento.VeiculoAno, orcamento.VeiculoCor, orcamento.Observacao)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting into orcamentos")
		return err
	}

	for _, item := range orcamentoItems {
		_, err := tx.Exec(`
      INSERT INTO orcamento_itens (orcamento_id, auto_categoria_id, observacao, data_criacao, data_atualizacao)
      VALUES ($1, $2, $3, now(), now())
    `, orcamento.ID, item.AutoCategoriaID, item.Observacao)
		if err != nil {
			log.Error().Err(err).Msg("Error inserting into orcamento_itens")
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

func (r *PgOrcamentoRepository) GetOrcamento(ctx context.Context, orcamentoID int64) (*types.OrcamentoModel, error) {
	o := types.Orcamento{}
	err := r.db.GetContext(ctx, &o, `
		SELECT id, associado_nome, veiculo_marca, veiculo_nome, veiculo_ano, veiculo_cor, 
    COALESCE(observacao, '') as observacao, status, data_criacao, data_atualizacao
		FROM orcamentos
		WHERE id = $1`, orcamentoID)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting orcamentos")
		return nil, err
	}

	var oi []types.OrcamentoItemModel
	err = r.db.SelectContext(ctx, &oi, `
    SELECT oi.id, oi.status, oi.data_criacao, oi.data_atualizacao, COALESCE(oi.observacao, '') as observacao, 
    ac.descricao as auto_categoria_descricao, ac.tipo as auto_categoria_tipo
    FROM orcamento_itens oi
    LEFT JOIN auto_categorias ac ON oi.auto_categoria_id = ac.id
    WHERE oi.orcamento_id = $1
    `, o.ID)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting orcamento_itens")
		return nil, err
	}

	return &types.OrcamentoModel{
		Orcamento: o,
		Itens:     oi,
	}, nil
}
