package middlewares

import (
	"github.com/labstack/echo/v4"
)

func HasRole(next echo.HandlerFunc, roles ...string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
