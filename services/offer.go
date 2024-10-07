package services

import (
	"context"

	"github.com/gmtborges/orcamento-auto/repos"
	"github.com/gmtborges/orcamento-auto/types"
)

type OfferService struct {
	offerRepo repos.OfferRepository
}

func NewOfferService(offerRepo repos.OfferRepository) *OfferService {
	return &OfferService{offerRepo: offerRepo}
}

func (s *OfferService) GetOfferByBiddingItemID(
	ctx context.Context,
	biddingItemID int64,
) ([]types.OfferModel, error) {
	return s.offerRepo.GetOffersByBiddingItemID(ctx, biddingItemID)
}
