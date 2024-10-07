package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/services"
	"github.com/gmtborges/orcamento-auto/types"
	"github.com/gmtborges/orcamento-auto/views/login"
)

type AuthHandler struct {
	userSvc *services.UserService
}

func NewAuthHandler(userSvc *services.UserService) *AuthHandler {
	return &AuthHandler{userSvc: userSvc}
}

func (h *AuthHandler) Index(c echo.Context) error {
	return views.LoginIndex(views.LoginIndexViewModel{}).Render(c.Request().Context(), c.Response())
}

func (h *AuthHandler) Login(c echo.Context) error {
	email := c.FormValue("email")
	passwd := c.FormValue("password")

	user, err := h.userSvc.GetUserByEmail(c.Request().Context(), email)
	if err == sql.ErrNoRows {
		c.Response().WriteHeader(http.StatusBadRequest)
		return views.LoginIndex(
			views.LoginIndexViewModel{
				Email:    email,
				Warnings: []string{"E-mail e/ou senha incorretos"},
			}).Render(c.Request().Context(), c.Response())
	}
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return views.LoginIndex(
			views.LoginIndexViewModel{
				Email:  email,
				Errors: []string{"Erro ao realizar login. Tente novamente mais tarde."},
			}).Render(c.Request().Context(), c.Response())
	}

	isValid, err := services.VerifyPasswordHash(passwd, user.Password)
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

	err = h.userSvc.SetSession(c, user.CompanyID, user.ID, user.Roles)
	if err != nil {
		log.Error().Err(err).Msg("Failed to set session")
		c.Response().WriteHeader(http.StatusInternalServerError)
		return views.LoginIndex(
			views.LoginIndexViewModel{
				Email:  email,
				Errors: []string{"Erro ao realizar login. Tente novamente mais tarde."},
			}).Render(c.Request().Context(), c.Response())
	}

	if user.CompanyType == string(types.CompanyTypeAuto) {
		return c.Redirect(http.StatusSeeOther, "/propostas")
	}
	return c.Redirect(http.StatusSeeOther, "/orcamentos")
}

func (h *AuthHandler) Logout(c echo.Context) error {
	if err := h.userSvc.RemoveSession(c); err != nil {
		// TODO: Send session error to the app layout toast alerts
		c.String(http.StatusInternalServerError, "Error removing session")
	}
	c.Response().Writer.Header().Add("HX-Redirect", "/entrar")
	return c.NoContent(http.StatusNoContent)
}
