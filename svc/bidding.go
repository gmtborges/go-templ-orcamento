package svc

import "github.com/gmtborges/orcamento-auto/repo"

type BiddingService struct {
	biddingRepo repo.BiddingRepository
}

func NewBiddingService(biddingRepo repo.BiddingRepository) *BiddingService {
	return &BiddingService{biddingRepo: biddingRepo}
}
