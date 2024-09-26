package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/views"
)

type PoliticaPrivacidadeHandler struct{}

func NewPoliticaPrivacidadeHandler() *PoliticaPrivacidadeHandler {
	return &PoliticaPrivacidadeHandler{}
}

func (h *PoliticaPrivacidadeHandler) Index(c echo.Context) error {
	return views.PoliticaPrivacidadeIndex().Render(c.Request().Context(), c.Response())
}
