package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/svc"
	"github.com/gmtborges/orcamento-auto/types"
	"github.com/gmtborges/orcamento-auto/views/proposta/partials"
)

type PropostaHandler struct {
	propostaSvc *svc.PropostaService
}

func NewPropostaHandler(propostaSvc *svc.PropostaService) *PropostaHandler {
	return &PropostaHandler{propostaSvc: propostaSvc}
}

func (h *PropostaHandler) GetPropostasByOrcamentoItemID(c echo.Context) error {
	oiID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NotFoundHandler(c)
	}
	p, err := h.propostaSvc.GetPropostaByOrcamentoItemID(c.Request().Context(), int64(oiID))
	vm := types.PropostaOrcamentoItensViewModel{
		Propostas: p,
	}

	return partials.PropostaList(vm).Render(c.Request().Context(), c.Response())
}
