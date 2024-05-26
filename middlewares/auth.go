package middlewares

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("auth-session", c)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		auth, ok := sess.Values["authenticated"].(bool)
		if !ok || !auth {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		return next(c)
	}
}
