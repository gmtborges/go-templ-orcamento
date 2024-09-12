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

func TestUserService_SetAuthSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	svc := NewUserService(&repo.MockUserRepository{})
	err := svc.SetSession(c, 321, 123, []string{"admin"})
	if assert.NoError(t, err) {
		s, err := session.Get("auth-session", c)
		if assert.NoError(t, err) {
			assert.Equal(t, true, s.Values["authenticated"])
			assert.Equal(t, int64(321), s.Values["company_id"])
			assert.Equal(t, int64(123), s.Values["user_id"])
			assert.Equal(t, []string{"admin"}, s.Values["roles"])
		}
	}
}

func TestUserService_RemoveAuthSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	svc := NewUserService(&repo.MockUserRepository{})
	if err := svc.SetSession(c, 321, 123, []string{"admin"}); err != nil {
		t.Errorf("Error setting session %v", err)
	}
	err := svc.RemoveUserSession(c)
	if assert.NoError(t, err) {
		s, err := session.Get("auth-session", c)
		if assert.NoError(t, err) {
			assert.Equal(t, false, s.Values["authenticated"], "Expected authenticated to be false")
			assert.Nil(t, s.Values["user_id"], "Expected user_id to be <nil>")
		}
	}
}
