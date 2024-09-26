package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/auth"
	"github.com/gmtborges/orcamento-auto/repo"
	"github.com/gmtborges/orcamento-auto/svc"
	"github.com/gmtborges/orcamento-auto/types"
)

func TestAutenticacaoHandler_Success(t *testing.T) {
	e := echo.New()
	hash, _ := auth.GeneratePasswordHash("passwd123")
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("email=test@example.com&password=passwd123"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))
	repo := &repo.MockUsuarioRepository{
		MockFn: func() (interface{}, error) {
			return types.UsuarioAutenticacao{
				ID:      int64(1),
				Nome:    "Tiao",
				Senha:   hash,
				Funcoes: []string{"admin"},
			}, nil
		},
	}
	usuarioSvc := svc.NewUsuarioService(repo)
	h := NewAutenticacaoHandler(usuarioSvc)

	if err := h.Login(c); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if rec.Code != http.StatusSeeOther {
		t.Errorf("Expected status See Other, got %v", rec.Code)
	}
}

func TestAutenticacaoHandler_usuarioNotFound(t *testing.T) {
	e := echo.New()
	repo := &repo.MockUsuarioRepository{
		MockFn: func() (interface{}, error) {
			return nil, sql.ErrNoRows
		},
	}
	usuarioSvc := svc.NewUsuarioService(repo)
	h := NewAutenticacaoHandler(usuarioSvc)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("email=test@example.com&password=passwd123"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.Login(c); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status Bad Request, got %v", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "E-mail e/ou senha incorretos") {
		t.Errorf("Expected body to contain E-mail e/ou senha incorretos, got %v", rec.Body.String())
	}
}

func TestAutenticacaoHandler_InvalidPassword(t *testing.T) {
	e := echo.New()
	repo := &repo.MockUsuarioRepository{
		MockFn: func() (interface{}, error) {
			return types.UsuarioAutenticacao{Senha: "wrong"}, nil
		},
	}
	usuarioSvc := svc.NewUsuarioService(repo)
	h := NewAutenticacaoHandler(usuarioSvc)
	req := httptest.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader("email=test@example.com&password=passwd123"),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.Login(c); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status Bad Request, got %v", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "E-mail e/ou senha incorretos") {
		t.Errorf("Expected body to contain E-mail e/ou senha incorretos, got %v", rec.Body.String())
	}
}

func TestAutenticacaoHandler_ErrorDatabase(t *testing.T) {
	e := echo.New()
	repo := &repo.MockUsuarioRepository{
		MockFn: func() (interface{}, error) {
			return nil, errors.New("error on database")
		},
	}
	usuarioSvc := svc.NewUsuarioService(repo)
	h := NewAutenticacaoHandler(usuarioSvc)
	req := httptest.NewRequest(
		http.MethodPost,
		"/entrar",
		strings.NewReader("email=test@example.com&password=passwd123"),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.Login(c); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status Internal Server Error, got %v", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Erro ao realizar login.") {
		t.Errorf("Expected body to contain Erro ao realizar login., got %v", rec.Body.String())
	}
}
