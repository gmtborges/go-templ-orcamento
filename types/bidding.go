package types

import (
	"time"
)

type BiddingStatus string

const (
	BiddingStatusAwaitingOffer BiddingStatus = "AWAITING_OFFER"
	BiddingStatusPending       BiddingStatus = "PENDING"
	BiddingStatusCanceled      BiddingStatus = "CANCELED"
	BiddingStatusFinished      BiddingStatus = "FINISHED"
)

type BiddingItemStatus string

const (
	BiddingItemStatusOpen          BiddingItemStatus = "OPEN"
	BiddingItemStatusOfferReceived BiddingItemStatus = "OFFER_RECEIVED"
	BiddingItemStatusOfferAccepted BiddingItemStatus = "OFFER_ACCEPTED"
	BiddingItemStatusCanceled      BiddingItemStatus = "CANCELED"
)

type Bidding struct {
	ID           int64
	CompanyID    int64  `db:"company_id"`
	CompanyPhone string `db:"company_phone"`
	UserID       int64  `db:"user_id"`
	CustomerName string `db:"customer_name" form:"customerName"`
	VehicleBrand string `db:"vehicle_brand" form:"vehicleBrand"`
	VehicleName  string `db:"vehicle_name" form:"vehicleName"`
	VehicleYear  int    `db:"vehicle_year" form:"vehicleYear"`
	VehicleColor string `db:"vehicle_color" form:"vehicleColor"`
	Notes        string `form:"notes"`
	Status       BiddingStatus
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type BiddingItem struct {
	ID                      int64
	BiddingID               int64 `db:"bidding_id"`
	Notes                   string
	Status                  BiddingItemStatus
	AutoCategoryID          int64     `db:"auto_category_id"`
	AutoCategoryDescription string    `db:"auto_category_description"`
	AutoCategoryType        string    `db:"auto_category_type"`
	CreatedAt               time.Time `db:"created_at"`
	UpdatedAt               time.Time `db:"updated_at"`
}

type BiddingItemOffers struct {
	BiddingItem
	Offers []Offer
}

type BiddingBiddingItems struct {
	Bidding
	Items []BiddingItem
}

type BiddingBiddingItemsOffers struct {
	Bidding
	Items []BiddingItemOffers
}

type BiddingFilters struct {
	Limit      int
	Offset     int
	OrderBy    string
	Order      string
	FilterBy   string
	SearchTerm string
}

type BiddingResultSet struct {
	Count int
	Data  []BiddingBiddingItems
}

type BiddingAutoResultSet struct {
	Count int
	Data  []BiddingBiddingItemsOffers
}

type BiddingIndexViewModel struct {
	Count       int
	CurrentPage int
	TotalPages  int
	SeqNumber   int
	Biddings    []BiddingBiddingItems
	Errors      []string
}

type BiddingNewViewModel struct {
	BiddingBiddingItems
	AutoCategories map[string][]AutoCategory
	Errors         map[string]string
}

type BiddingShowViewModel struct {
	BiddingBiddingItems
	Errors map[string]string
}

type BiddingItemOffersViewModel struct {
	Offers []Offer
	Errors map[string]string
}
