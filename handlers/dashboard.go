package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/services"
)

type DashHandler struct {
	DashSvc *services.DashService
}

func NewDashHandler(dashSvc *services.DashService) *DashHandler {
	return &DashHandler{DashSvc: dashSvc}
}

func (h *DashHandler) Show(c echo.Context) error {
	return c.String(http.StatusOK, "Dashboard")
}
