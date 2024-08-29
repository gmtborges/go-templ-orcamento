package types

import "time"

type Bidding struct {
	ID          int64
	Description string
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type BiddingSlice []*Bidding
