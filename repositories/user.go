package repositories

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/gustavomtborges/orcamento-auto/models"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (st *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := models.Users(qm.Where("email = ?", email)).One(ctx, st.db)
	if err != nil {
		return nil, err
	}
	return user, nil
}
