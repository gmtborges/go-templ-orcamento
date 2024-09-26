package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/gmtborges/orcamento-auto/auth"
	"github.com/gmtborges/orcamento-auto/svc"
	"github.com/gmtborges/orcamento-auto/views/login"
)

type AutenticacaoHandler struct {
	usuarioSvc *svc.UsuarioService
}

func NewAutenticacaoHandler(userSvc *svc.UsuarioService) *AutenticacaoHandler {
	return &AutenticacaoHandler{usuarioSvc: userSvc}
}

func (h *AutenticacaoHandler) Index(c echo.Context) error {
	return views.LoginIndex(views.LoginIndexViewModel{}).Render(c.Request().Context(), c.Response())
}

func (h *AutenticacaoHandler) Login(c echo.Context) error {
	email := c.FormValue("email")
	passwd := c.FormValue("password")

	usuario, err := h.usuarioSvc.GetUserByEmail(c.Request().Context(), email)
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

	isValid, err := auth.VerifyPasswordHash(passwd, usuario.Senha)
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

	err = h.usuarioSvc.SetSession(c, usuario.EmpresaID, usuario.ID, usuario.Funcoes)
	if err != nil {
		log.Error().Err(err).Msg("Failed to set session")
		c.Response().WriteHeader(http.StatusInternalServerError)
		return views.LoginIndex(
			views.LoginIndexViewModel{
				Email:  email,
				Errors: []string{"Erro ao realizar login. Tente novamente mais tarde."},
			}).Render(c.Request().Context(), c.Response())
	}

	return c.Redirect(http.StatusSeeOther, "/orcamentos")
}

func (h *AutenticacaoHandler) Logout(c echo.Context) error {
	if err := h.usuarioSvc.RemoveSession(c); err != nil {
		// TODO: Send session error to the app layout toast alerts
		c.String(http.StatusInternalServerError, "Error removing session")
	}
	c.Response().Writer.Header().Add("HX-Redirect", "/entrar")
	return c.NoContent(http.StatusNoContent)
}
