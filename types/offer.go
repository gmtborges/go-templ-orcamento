package types

import "time"

type Offer struct {
	ID            int64
	BiddingItemID int64  `db:"bidding_item_id"`
	CompanyID     int64  `db:"company_id"`
	CompanyName   string `db:"company_name"`
	Notes         string
	Price         float32
	Accepted      bool
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type OfferIndexViewModel struct {
	Count       int
	CurrentPage int
	TotalPages  int
	SeqNumber   int
	Biddings    []BiddingBiddingItemsOffers
	Errors      []string
}
