package repos

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/types"
)

type UserRepository interface {
	GetByID(ctx context.Context, uID int64) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*types.UserAuth, error)
}

type PgUserRepository struct {
	db *sqlx.DB
}

func NewPgUserRepository(db *sqlx.DB) *PgUserRepository {
	return &PgUserRepository{db: db}
}

func (r *PgUserRepository) GetUserByEmail(ctx context.Context, email string) (*types.UserAuth, error) {
	u := types.UserAuth{}
	err := r.db.GetContext(ctx, &u, `SELECT u.id, c.id as company_id, c.type as company_type, u.name, u.password
  FROM users u 
  LEFT JOIN companies c ON u.company_id = c.id
  WHERE email = $1 LIMIT 1`, email)
	if err != nil {
		log.Error().Err(err).Msg("Error getting user")
		return nil, err
	}
	roles := []string{}
	err = r.db.SelectContext(ctx, &roles, `SELECT r.name FROM roles r
  LEFT JOIN users_roles ur ON ur.role_id = r.id 
  WHERE ur.user_id = $1`, u.ID)
	if err != nil {
		log.Error().Err(err).Msg("Error selecting roles")
		return nil, err
	}
	u.Roles = roles

	return &u, nil
}

func (r *PgUserRepository) GetByID(ctx context.Context, userID int64) (int64, error) {
	var id int64
	err := r.db.GetContext(ctx, &id, "SELECT id FROM users WHERE id = $1", userID)
	if err != nil {
		return 0, err
	}
	return id, nil
}
