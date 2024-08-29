package types

import "time"

type Company struct {
	ID            int64
	Name          string
	Address       string
	ContactNumber string    `db:"contact_number"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
