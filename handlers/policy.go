package handlers

import (
	"github.com/labstack/echo/v4"

	"orcamento-auto/views"
)

func HandleGetPolicy(c echo.Context) error {
	return renderer(c, views.Policy())
}
