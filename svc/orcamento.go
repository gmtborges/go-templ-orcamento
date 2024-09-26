package svc

import (
	"context"

	"github.com/gmtborges/orcamento-auto/repo"
	"github.com/gmtborges/orcamento-auto/types"
)

type OrcamentoService struct {
	orcamentoRepo     repo.OrcamentoRepository
	autoCategoriaRepo repo.AutoCategoriaRepository
}

func NewOrcamentoService(
	orcamentoRepo repo.OrcamentoRepository,
	autoCategoriaRepo repo.AutoCategoriaRepository,
) *OrcamentoService {
	return &OrcamentoService{orcamentoRepo: orcamentoRepo, autoCategoriaRepo: autoCategoriaRepo}
}

func (s *OrcamentoService) GetAllOrcamentos(
	ctx context.Context,
	empID int64,
	filtros types.OrcamentoFiltros,
) (*types.OrcamentoResultSet, error) {
	return s.orcamentoRepo.GetAllOrcamentos(ctx, empID, filtros)
}

func (s *OrcamentoService) CreateOrcamento(
	ctx context.Context,
	uID, empID int64,
	orcamento types.Orcamento,
	orcamentoItens []struct{ types.OrcamentoItem },
) error {
	return s.orcamentoRepo.CreateOrcamento(ctx, uID, empID, orcamento, orcamentoItens)
}

func (s *OrcamentoService) GetOrcamento(ctx context.Context, orcamentoID int64) (*types.OrcamentoModel, error) {
	return s.orcamentoRepo.GetOrcamento(ctx, orcamentoID)
}

func (s *OrcamentoService) GetAutoCategorias(ctx context.Context) (map[string][]types.AutoCategoria, error) {
	acs, err := s.autoCategoriaRepo.GetAllAutoCategorias(ctx)
	if err != nil {
		return nil, err
	}
	var acProduto []types.AutoCategoria
	var acServico []types.AutoCategoria
	for _, ac := range acs {
		if ac.Tipo == types.AutoCategoriaTipoProduto {
			acProduto = append(acProduto, ac)
		}
		if ac.Tipo == types.AutoCategoriaTipoServico {
			acServico = append(acServico, ac)
		}
	}
	acGroup := map[string][]types.AutoCategoria{"acProduto": acProduto, "acServico": acServico}
	return acGroup, nil
}

func SendWhatsapp() {
	// Zap Orcamento Auto +55 62 9 9667-2684
}
