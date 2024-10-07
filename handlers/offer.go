package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/services"
	"github.com/gmtborges/orcamento-auto/types"
	"github.com/gmtborges/orcamento-auto/views/bidding/partials"
	views "github.com/gmtborges/orcamento-auto/views/offer"
)

type OfferHandler struct {
	offerSvc *services.OfferService
}

func NewOfferHandler(offerSvc *services.OfferService) *OfferHandler {
	return &OfferHandler{offerSvc: offerSvc}
}

func (h *OfferHandler) Index(c echo.Context) error {
	vm := types.OfferIndexViewModel{}
	return views.OfferIndex(vm).Render(c.Request().Context(), c.Response())
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
