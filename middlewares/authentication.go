package middlewares

import (
	"errors"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/services"
)

type Session struct {
	UserID    int64
	CompanyID int64
	Roles     []string
}

func Authentication(userSvc *services.UserService) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if s, err := getSession(c); err != nil {
				return c.Redirect(http.StatusSeeOther, "/entrar")
			} else {
				_, err := userSvc.GetByID(c.Request().Context(), s.UserID)
				if err != nil {
					return c.Redirect(http.StatusSeeOther, "/entrar")
				}
				c.Set("userID", s.UserID)
				c.Set("companyID", s.CompanyID)
				c.Set("roles", s.Roles)
				return next(c)
			}
		}
	}
}

func getSession(c echo.Context) (*Session, error) {
	sess, err := session.Get("session", c)
	if err != nil {
		return nil, err
	}
	userID, ok := sess.Values["userID"].(int64)
	if !ok {
		return nil, errors.New("'userID' was not found in the session.")
	}
	companyID, ok := sess.Values["companyID"].(int64)
	if !ok {
		return nil, errors.New("'companyID' was not found in the session.")
	}
	roles, ok := sess.Values["roles"].([]string)
	if !ok {
		return nil, errors.New("'roles' were not found in the session.")
	}
	return &Session{UserID: userID, CompanyID: companyID, Roles: roles}, nil
}
