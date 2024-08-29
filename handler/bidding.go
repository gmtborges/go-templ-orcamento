package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/svc"
	"github.com/gmtborges/orcamento-auto/views/bidding"
)

type BiddingHandler struct {
	BiddingSvc *svc.BiddingService
}

func NewBiddingsHandler(biddingSvc *svc.BiddingService) *BiddingHandler {
	return &BiddingHandler{BiddingSvc: biddingSvc}
}

func (h *BiddingHandler) Index(c echo.Context) error {
	return views.BiddingIndex().Render(c.Request().Context(), c.Response())
}
