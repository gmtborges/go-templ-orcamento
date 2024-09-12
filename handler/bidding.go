package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/svc"
	"github.com/gmtborges/orcamento-auto/types"
	views "github.com/gmtborges/orcamento-auto/views/bidding"
	"github.com/gmtborges/orcamento-auto/views/bidding/partials"
)

type BiddingHandler struct {
	BiddingSvc *svc.BiddingService
}

func NewBiddingsHandler(biddingSvc *svc.BiddingService) *BiddingHandler {
	return &BiddingHandler{BiddingSvc: biddingSvc}
}

func (h *BiddingHandler) Index(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("p"))
	searchTerm := c.QueryParam("q")
	if err != nil {
		page = 1
	}
	offset := (page - 1) * 10
	filters := types.BiddingFilters{
		Limit:      10,
		Offset:     offset,
		OrderBy:    "created_at",
		Order:      "asc",
		FilterBy:   "",
		SearchTerm: searchTerm,
	}
	companyID := c.Get("company_id").(int64)
	isHTMX := c.Request().Header.Get("HX-Request") == "true"
	result, err := h.BiddingSvc.AllBiddings(c.Request().Context(), companyID, filters)
	if err != nil {
		vm := types.BiddingIndexViewModel{
			Errors: []string{"Erro ao buscar os orçamentos. Tente novamente mais tarde."},
		}
		return views.BiddingIndex(vm).Render(c.Request().Context(), c.Response())
	}
	vm := types.BiddingIndexViewModel{
		Count:       result.Count,
		CurrentPage: page,
		TotalPages:  (result.Count + filters.Limit - 1) / filters.Limit,
		SeqNumber:   (page-1)*filters.Limit + 1,
		Biddings:    result.Data,
		Errors:      []string{},
	}
	if isHTMX {
		return partials.BiddingList(vm).Render(c.Request().Context(), c.Response())
	}
	return views.BiddingIndex(vm).Render(c.Request().Context(), c.Response())
}

func (h *BiddingHandler) New(c echo.Context) error {
	return views.BiddingNew().Render(c.Request().Context(), c.Response())
}

// func (h *BiddingHandler) Show(c echo.Context) error {
// 	companyID := c.Get("company_id").(int64)
// 	biddingID := c.Param("id")
// 	vm, err := h.BiddingSvc.GetBidding(c.Request().Context(), companyID, biddingID)
// 	if err != nil {
// 		vm.Errors = []string{"Erro ao buscar o orçamento. Tente novamente mais tarde."}
// 		return views.BiddingShow(*vm).Render(c.Request().Context(), c.Response())
// 	}
// 	return views.BiddingShow(*vm).Render(c.Request().Context(), c.Response())
// }
