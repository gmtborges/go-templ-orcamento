package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/types"
)

type UserRepository interface {
	GetByID(ctx context.Context, userID int64) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*types.UserAuth, error)
}

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*types.UserAuth, error) {
	u := types.UserAuth{}
	err := r.db.GetContext(ctx, &u, "SELECT id, company_id, name, password FROM users WHERE email = $1 LIMIT 1", email)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user by email")
		return nil, err
	}
	rs := []string{}
	err = r.db.SelectContext(ctx, &rs, `SELECT r.name FROM roles r
  LEFT JOIN users_roles ur ON ur.role_id = r.id 
  WHERE ur.user_id = $1`, u.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user roles")
		return nil, err
	}
	u.Roles = rs

	return &u, nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, userID int64) (int64, error) {
	var id int64
	err := r.db.GetContext(ctx, &id, "SELECT id FROM users WHERE id = $1", userID)
	if err != nil {
		return 0, err
	}
	return id, nil
}
