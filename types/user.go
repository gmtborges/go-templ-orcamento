package types

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Email     string
	Password  string
	Name      string
	CompanyID sql.NullInt64 `db:"company_id"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt time.Time     `db:"updated_at"`
}

type Role struct {
	ID   int64
	Name string
}

type UserAuth struct {
	ID        int64
	CompanyID int64 `db:"company_id"`
	Name      string
	Password  string
	Roles     []string
}
