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
		userID, ok := sess.Values["user_id"].(string)
		if !ok {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		c.Set("user_id", userID)
		return next(c)
	}
}
