package stores

import (
	"context"
	"database/sql"

	"github.com/gustavomtborges/orcamento-auto/models"
)

type CompanyStorer interface {
	GetCompany(ctx context.Context, id int) (*models.Company, error)
}

type PostgresCompanyStore struct {
	db *sql.DB
}

func NewPostgresCompanyStore(db *sql.DB) *PostgresCompanyStore {
	return &PostgresCompanyStore{db: db}
}

func (st *PostgresCompanyStore) GetCompany(ctx context.Context, id int) (*models.Company, error) {
	company, err := models.Companies(models.CompanyWhere.ID.EQ(id)).One(ctx, st.db)
	if err != nil {
		return nil, err
	}
	return company, nil
}
