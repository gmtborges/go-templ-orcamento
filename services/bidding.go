package services

import "github.com/gmtborges/orcamento-auto/repositories"

type BiddingService struct {
	biddingRepo repositories.BiddingRepository
}

func NewBiddingService(biddingRepo repositories.BiddingRepository) *BiddingService {
	return &BiddingService{biddingRepo: biddingRepo}
}
