package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"

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
	u := &types.UserAuth{}
	err := r.db.GetContext(ctx, &u, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, userID int64) (int64, error) {
	var id int64
	err := r.db.GetContext(ctx, &id, "SELECT id FROM users WHERE id = $1", userID)
	if err != nil {
		return 0, err
	}
	return id, nil
}
