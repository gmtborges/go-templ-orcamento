package stores

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/gustavomtborges/orcamento-auto/models"
)

type UserStorer interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{db: db}
}

func (st *PostgresUserStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := models.Users(qm.Where("email = ?", email)).One(ctx, st.db)
	if err != nil {
		return nil, err
	}
	return user, nil
}
