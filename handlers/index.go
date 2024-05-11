package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/views"
)

func ShowIndex(c echo.Context) error {
	return render(c, views.Index())
}
