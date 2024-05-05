package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/views"
)

func HandleGetIndex(c echo.Context) error {
	return renderer(c, views.Index())
}
