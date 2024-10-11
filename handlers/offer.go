package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/services"
	"github.com/gmtborges/orcamento-auto/types"
	"github.com/gmtborges/orcamento-auto/views/pages"
	"github.com/gmtborges/orcamento-auto/views/partials"
)

type OfferHandler struct {
	offerSvc   *services.OfferService
	biddingSvc *services.BiddingService
}

func NewOfferHandler(offerSvc *services.OfferService, biddingSvc *services.BiddingService) *OfferHandler {
	return &OfferHandler{
		offerSvc:   offerSvc,
		biddingSvc: biddingSvc,
	}
}

func (h *OfferHandler) Index(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("p"))
	searchTerm := c.QueryParam("q")
	if err != nil {
		page = 1
	}
	offset := (page - 1) * 3
	filters := types.BiddingFilters{
		Limit:      3,
		Offset:     offset,
		OrderBy:    "created_at",
		Order:      "desc",
		FilterBy:   "",
		SearchTerm: searchTerm,
	}
	companyID := c.Get("companyID").(int64)
	isHTMX := c.Request().Header.Get("HX-Request") == "true"
	result, err := h.biddingSvc.GetAllBiddingsByAutoCategoryIDs(c.Request().Context(), companyID, filters)
	if err != nil {
		vm := types.OfferIndexViewModel{
			Errors: []string{"Erro ao buscar as propostas. Tente novamente mais tarde."},
		}
		return pages.OfferIndex(vm).Render(c.Request().Context(), c.Response())
	}
	vm := types.OfferIndexViewModel{
		Count:       result.Count,
		CurrentPage: page,
		TotalPages:  (result.Count + filters.Limit - 1) / filters.Limit,
		Biddings:    result.Data,
		Errors:      []string{},
	}
	if isHTMX {
		return partials.OfferList(vm).Render(c.Request().Context(), c.Response())
	}
	return pages.OfferIndex(vm).Render(c.Request().Context(), c.Response())
}

func (h *OfferHandler) GetOffersByBiddingItemID(c echo.Context) error {
	biddingItemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NotFoundHandler(c)
	}
	o, err := h.offerSvc.GetOfferByBiddingItemID(c.Request().Context(), int64(biddingItemID))
	vm := types.BiddingItemOffersViewModel{
		Offers: o,
	}

	return partials.BiddingItemOfferList(vm).Render(c.Request().Context(), c.Response())
}
