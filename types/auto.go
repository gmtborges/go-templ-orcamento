package types

import "time"

type AutoOffer struct {
	ID           int64     `db:"id"`
	BiddingID    int64     `db:"bidding_id"`
	OfferDetails string    `db:"offer_details"`
	OfferDate    time.Time `db:"offer_date"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type CategoryType string

const (
	CategoryTypeProduct CategoryType = "PRODUCT"
	CategoryTypeService CategoryType = "SERVICE"
)

type AutoCategory struct {
	ID   int64
	Name string
	Type CategoryType
}
