package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/views"
)

type PolicyHandler struct{}

func NewPolicyHandler() *PolicyHandler {
	return &PolicyHandler{}
}

func (h *PolicyHandler) Index(c echo.Context) error {
	return views.Policy().Render(c.Request().Context(), c.Response())
}
