package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/services"
	"github.com/gmtborges/orcamento-auto/views/bidding"
)

type BiddingHandler struct {
	BiddingSvc *services.BiddingService
}

func NewBiddingsHandler(biddingSvc *services.BiddingService) *BiddingHandler {
	return &BiddingHandler{BiddingSvc: biddingSvc}
}

func (h *BiddingHandler) Index(c echo.Context) error {
	return views.BiddingIndex().Render(c.Request().Context(), c.Response())
}
