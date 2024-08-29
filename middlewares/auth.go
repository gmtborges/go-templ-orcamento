package middlewares

import (
	"errors"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/services"
)

func Authentication(userSvc *services.UserService) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if userID, roles, err := getSessionValues(c); err != nil {
				return c.Redirect(http.StatusSeeOther, "/login")
			} else {
				_, err := userSvc.GetByID(c.Request().Context(), userID)
				if err != nil {
					return c.Redirect(http.StatusSeeOther, "/login")
				}
				c.Set("user_id", userID)
				c.Set("roles", roles)
				return next(c)
			}
		}
	}
}

func getSessionValues(c echo.Context) (int64, string, error) {
	sess, err := session.Get("auth-session", c)
	if err != nil {
		return 0, "", err
	}
	userID, ok := sess.Values["user_id"].(int64)
	if !ok {
		return 0, "", errors.New("'user_id' was not found in the session.")
	}
	roles, ok := sess.Values["roles"].(string)
	if !ok {
		return 0, "", errors.New("'roles' were not found in the session.")
	}
	return userID, roles, nil
}
