package svc

import (
	"context"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/repo"
	"github.com/gmtborges/orcamento-auto/types"
)

type UsuarioService struct {
	usuarioRepo repo.UsuarioRepository
}

func NewUsuarioService(usuarioRepo repo.UsuarioRepository) *UsuarioService {
	return &UsuarioService{usuarioRepo: usuarioRepo}
}

func (s *UsuarioService) GetUserByEmail(ctx context.Context, email string) (*types.UsuarioAutenticacao, error) {
	usuario, err := s.usuarioRepo.GetUsuarioByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &types.UsuarioAutenticacao{
		ID:        usuario.ID,
		EmpresaID: usuario.EmpresaID,
		Nome:      usuario.Nome,
		Senha:     usuario.Senha,
		Funcoes:   usuario.Funcoes,
	}, nil
}

func (s *UsuarioService) GetByID(ctx context.Context, uID int64) (int64, error) {
	return s.usuarioRepo.GetByID(ctx, uID)
}

func (s *UsuarioService) SetSession(c echo.Context, empID int64, uID int64, funcoes []string) error {
	session, err := session.Get("sessao", c)
	if err != nil {
		return err
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	session.Values["uID"] = uID
	session.Values["empID"] = empID
	session.Values["funcoes"] = funcoes
	return session.Save(c.Request(), c.Response())
}

func (s *UsuarioService) IsAuthenticated(c echo.Context) (bool, error) {
	session, err := session.Get("sessao", c)
	if err != nil {
		return false, err
	}
	uID, ok := session.Values["uID"]
	return uID != nil && ok, nil
}

func (s *UsuarioService) RemoveSession(c echo.Context) error {
	session, err := session.Get("sessao", c)
	if err != nil {
		return err
	}
	session.Values["uID"] = nil
	session.Values["empID"] = nil
	session.Values["funcoes"] = nil
	session.Options.MaxAge = -1
	return session.Save(c.Request(), c.Response())
}
