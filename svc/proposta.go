package svc

import (
	"context"

	"github.com/gmtborges/orcamento-auto/repo"
	"github.com/gmtborges/orcamento-auto/types"
)

type PropostaService struct {
	propostaRepo repo.PropostaRepository
}

func NewPropostaService(propostaRepo repo.PropostaRepository) *PropostaService {
	return &PropostaService{propostaRepo: propostaRepo}
}

func (s *PropostaService) GetPropostaByOrcamentoItemID(
	ctx context.Context,
	oiID int64,
) ([]types.PropostaModel, error) {
	return s.propostaRepo.GetPropostasByOrcamentoItemID(ctx, oiID)
}
