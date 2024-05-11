package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/views"
)

func ShowLogin(c echo.Context) error {
	return render(c, views.Login())
}

func PostLogin(c echo.Context, db *sql.DB) error {
	email := c.FormValue("email")
	passwd := c.FormValue("password")

	return c.String(http.StatusOK, "email:"+email+", passwd:"+passwd)
}
