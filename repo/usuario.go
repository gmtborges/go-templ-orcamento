package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/types"
)

type UsuarioRepository interface {
	GetByID(ctx context.Context, uID int64) (int64, error)
	GetUsuarioByEmail(ctx context.Context, email string) (*types.UsuarioAutenticacao, error)
}

type PostgresUsuarioRepository struct {
	db *sqlx.DB
}

func NewPostgresUsuarioRepository(db *sqlx.DB) *PostgresUsuarioRepository {
	return &PostgresUsuarioRepository{db: db}
}

func (r *PostgresUsuarioRepository) GetUsuarioByEmail(ctx context.Context, email string) (*types.UsuarioAutenticacao, error) {
	u := types.UsuarioAutenticacao{}
	err := r.db.GetContext(ctx, &u, `SELECT id, empresa_id, nome, senha 
  FROM usuarios 
  WHERE email = $1 LIMIT 1`, email)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user by email")
		return nil, err
	}
	fs := []string{}
	err = r.db.SelectContext(ctx, &fs, `SELECT f.nome FROM funcoes f
  LEFT JOIN usuarios_funcoes uf ON uf.funcao_id = f.id 
  WHERE uf.usuario_id = $1`, u.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get usuarios_funcoes")
		return nil, err
	}
	u.Funcoes = fs

	return &u, nil
}

func (r *PostgresUsuarioRepository) GetByID(ctx context.Context, userID int64) (int64, error) {
	var id int64
	err := r.db.GetContext(ctx, &id, "SELECT id FROM usuarios WHERE id = $1", userID)
	if err != nil {
		return 0, err
	}
	return id, nil
}
