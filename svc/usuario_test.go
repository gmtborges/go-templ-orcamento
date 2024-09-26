package svc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/gmtborges/orcamento-auto/repo"
)

func TestUsuarioService_SetSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	svc := NewUsuarioService(&repo.MockUsuarioRepository{})
	err := svc.SetSession(c, 321, 123, []string{"admin"})
	if assert.NoError(t, err) {
		s, err := session.Get("sessao", c)
		if assert.NoError(t, err) {
			assert.Equal(t, int64(321), s.Values["empID"])
			assert.Equal(t, int64(123), s.Values["uID"])
			assert.Equal(t, []string{"admin"}, s.Values["funcoes"])
		}
	}
}

func TestUsuarioService_RemoveSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	svc := NewUsuarioService(&repo.MockUsuarioRepository{})
	if err := svc.SetSession(c, 321, 123, []string{"admin"}); err != nil {
		t.Errorf("Error setting session %v", err)
	}
	err := svc.RemoveSession(c)
	if assert.NoError(t, err) {
		s, err := session.Get("sessao", c)
		if assert.NoError(t, err) {
			assert.Nil(t, s.Values["uID"], "Expected uID to be <nil>")
		}
	}
}
