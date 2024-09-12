package svc

import (
	"context"

	"github.com/gmtborges/orcamento-auto/repo"
	"github.com/gmtborges/orcamento-auto/types"
)

type BiddingService struct {
	biddingRepo repo.BiddingRepository
}

func NewBiddingService(biddingRepo repo.BiddingRepository) *BiddingService {
	return &BiddingService{biddingRepo: biddingRepo}
}

func (s *BiddingService) AllBiddings(ctx context.Context, companyID int64, filters types.BiddingFilters) (*types.BiddingResult, error) {
	return s.biddingRepo.AllBiddings(ctx, companyID, filters)
}

func (s *BiddingService) CreateBidding(ctx context.Context, companyID int64, bidding *types.Bidding, biddingItems []types.BiddingItem) (*types.Bidding, error) {
	return s.biddingRepo.CreateBidding(ctx, companyID, bidding, biddingItems)
}
