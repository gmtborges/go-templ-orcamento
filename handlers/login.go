package handlers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/services"
	"github.com/gustavomtborges/orcamento-auto/views/login"
)

type LoginHandler struct {
	authSvc *services.AuthService
}

func NewLoginHandler(authSvc *services.AuthService) *LoginHandler {
	return &LoginHandler{authSvc: authSvc}
}

func (h *LoginHandler) Index(c echo.Context) error {
	return views.LoginIndex(views.LoginIndexViewModel{}).Render(c.Request().Context(), c.Response())
}

func (h *LoginHandler) Create(c echo.Context) error {
	email := c.FormValue("email")
	passwd := c.FormValue("password")

	user, err := h.authSvc.GetUserByEmail(c.Request().Context(), email)
	if err == sql.ErrNoRows {
		c.Response().WriteHeader(http.StatusBadRequest)
		return views.LoginIndex(
			views.LoginIndexViewModel{
				Email:    email,
				Warnings: []string{"E-mail e/ou senha incorretos"},
			}).Render(c.Request().Context(), c.Response())

	}
	if err != sql.ErrNoRows && err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return views.LoginIndex(
			views.LoginIndexViewModel{
				Email:  email,
				Errors: []string{"Erro ao realizar login. Tente novamente mais tarde."},
			}).Render(c.Request().Context(), c.Response())
	}

	isValid, err := h.authSvc.VerifyPasswordHash(passwd, user.Password)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		warnings := []string{"E-mail e/ou senha incorretos"}
		return views.LoginIndex(
			views.LoginIndexViewModel{
				Email:    email,
				Warnings: warnings,
			}).Render(c.Request().Context(), c.Response())
	}

	if !isValid {
		c.Response().WriteHeader(http.StatusUnauthorized)
		warnings := []string{"E-mail e/ou senha incorretos"}
		return views.LoginIndex(
			views.LoginIndexViewModel{
				Email:    email,
				Warnings: warnings,
			}).Render(c.Request().Context(), c.Response())
	}

	if err := h.authSvc.SetAuthSession(c, user.ID); err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return views.LoginIndex(
			views.LoginIndexViewModel{
				Email:  email,
				Errors: []string{"Erro ao realizar login. Tente novamente mais tarde."},
			}).Render(c.Request().Context(), c.Response())
	}

	return c.Redirect(http.StatusSeeOther, "/orcamentos")
}

func (h *LoginHandler) Logout(c echo.Context) error {
	if err := h.authSvc.RemoveAuthSession(c); err != nil {
		// TODO: Send session error to the app layout toast alerts
		c.String(http.StatusInternalServerError, "Error removing session")
	}
	c.Response().Writer.Header().Add("HX-Redirect", "/login")
	return c.NoContent(http.StatusNoContent)
}
