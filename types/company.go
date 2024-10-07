package types

import "time"

type CompanyType string

const (
	CompanyTypeOrg  CompanyType = "ORG"
	CompanyTypeAuto CompanyType = "AUTO"
)

type Company struct {
	ID        int64
	Name      string
	Type      CompanyType
	State     string
	City      string
	Address   string
	Cellphone string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
