package types

import (
	"database/sql"
	"time"
)

type BiddingStatus string

const (
	BiddingStatusAwaitingOffers BiddingStatus = "AWAITING_OFFERS"
	BiddingStatusPending        BiddingStatus = "PENDING"
	BiddingStatusCanceled       BiddingStatus = "CANCELED"
	BiddingStatusFinished       BiddingStatus = "FINISHED"
)

type BiddingItemStatus string

const (
	BiddingItemStatusOpen          BiddingItemStatus = "OPEN"
	BiddingItemStatusOfferReceived BiddingItemStatus = "OFFER_RECEIVED"
	BiddingItemStatusAccepted      BiddingItemStatus = "ACCEPTED"
	BiddingItemStatusCanceled      BiddingItemStatus = "CANCELED"
)

type Bidding struct {
	ID           int64
	CompanyID    int64  `db:"company_id"`
	UserID       int64  `db:"user_id"`
	CustomerName string `db:"customer_name"`
	VehicleBrand string `db:"vehicle_brand"`
	VehicleName  string `db:"vehicle_name"`
	VehicleYear  int    `db:"vehicle_year"`
	VehicleColor string `db:"vehicle_color"`
	Description  sql.NullString
	Status       BiddingStatus
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type BiddingItem struct {
	ID             int64
	BiddingID      int64 `db:"bidding_id"`
	Description    sql.NullString
	Status         BiddingItemStatus
	AutoCategoryID int64     `db:"auto_category_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type BiddingItemModel struct {
	BiddingItem
	AutoCategoryID   int64  `db:"auto_category_id"`
	AutoCategoryName string `db:"auto_category_name"`
	AutoCategoryType string `db:"auto_category_type"`
}

type BiddingModel struct {
	Bidding
	Items []BiddingItemModel
}

type BiddingResult struct {
	Count int
	Data  []BiddingModel
}

type BiddingIndexViewModel struct {
	Count       int
	CurrentPage int
	TotalPages  int
	SeqNumber   int
	Biddings    []BiddingModel
	Errors      []string
}

type BiddingShowViewModel struct {
	Bidding BiddingModel
	Errors  []string
}

type BiddingNewViewModel struct {
	Errors []string
}

type BiddingFilters struct {
	Limit      int
	Offset     int
	OrderBy    string
	Order      string
	FilterBy   string
	SearchTerm string
}
