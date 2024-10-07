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
	ID             int64
	BiddingID      int64 `db:"bidding_id"`
	Notes          string
	Status         BiddingItemStatus
	AutoCategoryID int64     `db:"auto_category_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type BiddingItemModel struct {
	BiddingItem
	AutoCategoryDescription string `db:"auto_category_description"`
	AutoCategoryType        string `db:"auto_category_type"`
}

type BiddingItemOfferModel struct {
	BiddingItem
	Offer
}

type BiddingModel struct {
	Bidding
	Items []BiddingItemModel
}

type BiddingAutoModel struct {
	Bidding
	Items []BiddingItemOfferModel
}

type BiddingResultSet struct {
	Count int
	Data  []BiddingModel
}

type BiddingAutoResultSet struct {
	Count int
	Data  []BiddingAutoModel
}

type BiddingIndexViewModel struct {
	Count       int
	CurrentPage int
	TotalPages  int
	SeqNumber   int
	Biddings    []BiddingModel
	Errors      []string
}

type BiddingCreateViewModel struct {
	BiddingModel
	AutoCategories map[string][]AutoCategory
	Errors         map[string]string
}

type BiddingShowViewModel struct {
	BiddingModel
	Errors map[string]string
}

type BiddingFilters struct {
	Limit      int
	Offset     int
	OrderBy    string
	Order      string
	FilterBy   string
	SearchTerm string
}

type BiddingItemOffersViewModel struct {
	Offers []OfferModel
	Errors map[string]string
}
