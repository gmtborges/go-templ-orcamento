package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/views/pages"
)

type PolicyHandler struct{}

func NewPolicyHandler() *PolicyHandler {
	return &PolicyHandler{}
}

func (h *PolicyHandler) Index(c echo.Context) error {
	return pages.PolicyIndex().Render(c.Request().Context(), c.Response())
}
