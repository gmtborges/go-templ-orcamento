package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/services"
	"github.com/gustavomtborges/orcamento-auto/views/bidding"
)

type BiddingHandler struct {
	BiddingSvc *services.BiddingService
}

func NewBiddingsHandler(dashSvc *services.BiddingService) *BiddingHandler {
	return &BiddingHandler{BiddingSvc: dashSvc}
}

func (h *BiddingHandler) Index(c echo.Context) error {
	fmt.Print("why not here?")
	return views.BiddingIndex().Render(c.Request().Context(), c.Response())
}
