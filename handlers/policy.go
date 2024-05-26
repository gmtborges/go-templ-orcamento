package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/views"
)

type PolicyHandler struct{}

func NewPolicyHandler() *PolicyHandler {
	return &PolicyHandler{}
}

func (h *PolicyHandler) Show(c echo.Context) error {
	return views.Policy().Render(c.Request().Context(), c.Response())
}
