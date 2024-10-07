package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/gmtborges/orcamento-auto/repos"
)

func TestUserService_SetSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	services := NewUserService(&repos.MockUserRepository{})
	err := services.SetSession(c, 321, 123, []string{"admin"})
	if assert.NoError(t, err) {
		s, err := session.Get("session", c)
		if assert.NoError(t, err) {
			assert.Equal(t, int64(321), s.Values["companyID"])
			assert.Equal(t, int64(123), s.Values["userID"])
			assert.Equal(t, []string{"admin"}, s.Values["roles"])
		}
	}
}

func TestUserService_RemoveSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	services := NewUserService(&repos.MockUserRepository{})
	if err := services.SetSession(c, 321, 123, []string{"admin"}); err != nil {
		t.Errorf("Error setting session %v", err)
	}
	err := services.RemoveSession(c)
	if assert.NoError(t, err) {
		s, err := session.Get("session", c)
		if assert.NoError(t, err) {
			assert.Nil(t, s.Values["userID"], "Expected userID to be <nil>")
		}
	}
}
