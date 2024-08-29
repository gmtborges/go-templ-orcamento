package types

import "time"

type User struct {
	ID        int64
	Email     string
	Password  string
	Name      string
	CompanyID int64     `db:"company_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserAuth struct {
	ID       int64
	Name     string
	Password string
	Roles    []string
}
