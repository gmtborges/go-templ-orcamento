package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func renderer(c echo.Context, cp templ.Component) error {
	return cp.Render(c.Request().Context(), c.Response())
}
