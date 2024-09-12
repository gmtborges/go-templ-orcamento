package types

import "time"

type CompanyType string

const (
	CompanyTypeOrg  CompanyType = "ORG"
	CompanyTypeAuto CompanyType = "AUTO"
)

type Company struct {
	ID            int64
	Name          string
	Type          CompanyType
	Province      string
	City          string
	Address       string
	ContactNumber string    `db:"contact_number"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
