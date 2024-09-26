package middlewares

import (
	"errors"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/svc"
)

type Sessao struct {
	UsuarioID int64
	EmpresaID int64
	Funcoes   []string
}

func Autenticacao(usuarioSvc *svc.UsuarioService) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if sessao, err := getSession(c); err != nil {
				return c.Redirect(http.StatusSeeOther, "/entrar")
			} else {
				_, err := usuarioSvc.GetByID(c.Request().Context(), sessao.UsuarioID)
				if err != nil {
					return c.Redirect(http.StatusSeeOther, "/entrar")
				}
				c.Set("uID", sessao.UsuarioID)
				c.Set("empID", sessao.EmpresaID)
				c.Set("funcoes", sessao.Funcoes)
				return next(c)
			}
		}
	}
}

func getSession(c echo.Context) (*Sessao, error) {
	sess, err := session.Get("sessao", c)
	if err != nil {
		return nil, err
	}
	uID, ok := sess.Values["uID"].(int64)
	if !ok {
		return nil, errors.New("'uID' was not found in the session.")
	}
	empID, ok := sess.Values["empID"].(int64)
	if !ok {
		return nil, errors.New("'empID' was not found in the session.")
	}
	funcoes, ok := sess.Values["funcoes"].([]string)
	if !ok {
		return nil, errors.New("'funcoes' were not found in the session.")
	}
	return &Sessao{UsuarioID: uID, EmpresaID: empID, Funcoes: funcoes}, nil
}
