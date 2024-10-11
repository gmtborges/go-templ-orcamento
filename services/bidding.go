package services

import (
	"context"

	"github.com/gmtborges/orcamento-auto/repos"
	"github.com/gmtborges/orcamento-auto/types"
)

type BiddingService struct {
	biddingRepo      repos.BiddingRepository
	autoCategoryRepo repos.AutoCategoryRepository
}

func NewBiddingService(
	biddingRepo repos.BiddingRepository,
	autoCategoryRepo repos.AutoCategoryRepository,
) *BiddingService {
	return &BiddingService{biddingRepo: biddingRepo, autoCategoryRepo: autoCategoryRepo}
}

func (s *BiddingService) GetAllBiddingsByCompanyID(
	ctx context.Context,
	companyID int64,
	filters types.BiddingFilters,
) (*types.BiddingResultSet, error) {
	return s.biddingRepo.GetAllBiddingsByCompanyID(ctx, companyID, filters)
}

func (s *BiddingService) GetAllBiddingsByAutoCategoryIDs(
	ctx context.Context,
	companyID int64,
	filters types.BiddingFilters,
) (*types.BiddingAutoResultSet, error) {
	acIDs, err := s.autoCategoryRepo.GetAllAutoCategoryIDsByCompanyID(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return s.biddingRepo.GetAllBiddingsByAutoCategoryIDs(ctx, acIDs, filters)
}

func (s *BiddingService) CreateBidding(
	ctx context.Context,
	userID, companyID int64,
	bidding types.Bidding,
	biddingItems []struct{ types.BiddingItem },
) error {
	return s.biddingRepo.CreateBidding(ctx, userID, companyID, bidding, biddingItems)
}

func (s *BiddingService) GetBidding(ctx context.Context, biddingID int64) (*types.BiddingBiddingItems, error) {
	return s.biddingRepo.GetBidding(ctx, biddingID)
}

func (s *BiddingService) GetAutoCategories(ctx context.Context) (map[string][]types.AutoCategory, error) {
	acs, err := s.autoCategoryRepo.GetAllAutoCategories(ctx)
	if err != nil {
		return nil, err
	}
	var acProduct []types.AutoCategory
	var acService []types.AutoCategory
	for _, ac := range acs {
		if ac.Type == types.AutoCategoryTypeProduct {
			acProduct = append(acProduct, ac)
		}
		if ac.Type == types.AutoCategoryTypeService {
			acService = append(acService, ac)
		}
	}
	acGroup := map[string][]types.AutoCategory{"acProduct": acProduct, "acService": acService}
	return acGroup, nil
}

func SendWhatsapp() {
	// Zap Orcamento Auto +55 62 9 9667-2684
}
