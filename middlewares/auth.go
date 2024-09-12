package middlewares

import (
	"errors"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/svc"
)

type SessionValues struct {
	UserID    int64
	CompanyID int64
	Roles     []string
}

func Authentication(userSvc *svc.UserService) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if sessionValues, err := getSessionValues(c); err != nil {
				return c.Redirect(http.StatusSeeOther, "/login")
			} else {
				_, err := userSvc.GetByID(c.Request().Context(), sessionValues.UserID)
				if err != nil {
					return c.Redirect(http.StatusSeeOther, "/login")
				}
				c.Set("user_id", sessionValues.UserID)
				c.Set("company_id", sessionValues.CompanyID)
				c.Set("roles", sessionValues.Roles)
				return next(c)
			}
		}
	}
}

func getSessionValues(c echo.Context) (*SessionValues, error) {
	sess, err := session.Get("auth-session", c)
	if err != nil {
		return nil, err
	}
	userID, ok := sess.Values["user_id"].(int64)
	if !ok {
		return nil, errors.New("'user_id' was not found in the session.")
	}
	companyID, ok := sess.Values["company_id"].(int64)
	if !ok {
		return nil, errors.New("'company_id' was not found in the session.")
	}
	roles, ok := sess.Values["roles"].([]string)
	if !ok {
		return nil, errors.New("'roles' were not found in the session.")
	}
	return &SessionValues{UserID: userID, CompanyID: companyID, Roles: roles}, nil
}
