package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/views"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) Show(c echo.Context) error {
	return views.Index().Render(c.Request().Context(), c.Response())
}
