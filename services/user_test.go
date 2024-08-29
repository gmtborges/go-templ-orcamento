package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/repositories"
)

func TestUserService_SetAuthSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	svc := NewUserService(&repositories.MockUserRepository{})
	err := svc.SetSession(c, 123, "assoc_admin,auto_admin")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	s, err := session.Get("auth-session", c)
	if err != nil {
		t.Errorf("Error getting session: %v", err)
	}
	if s.Values["authenticated"] != true {
		t.Errorf("Expected authenticated to be true, got %v", s.Values["authenticated"])
	}
	if s.Values["user_id"].(int64) != 123 {
		t.Errorf("Expected user_id to be 123, got %v", s.Values["user_id"])
	}
	if s.Values["roles"] != "assoc_admin,auto_admin" {
		t.Errorf("Expected user_id to be 123, got %v", s.Values["user_id"])
	}
}

func TestUserService_RemoveAuthSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	svc := NewUserService(&repositories.MockUserRepository{})
	if err := svc.SetSession(c, 123, "admin"); err != nil {
		t.Errorf("Error setting session %v", err)
	}
	err := svc.RemoveUserSession(c)
	if err != nil {
		t.Errorf("Error removing session %v", err)
	}
	s, err := session.Get("auth-session", c)
	if err != nil {
		t.Errorf("Error getting session %v", err)
	}
	if s.Values["authenticated"] != false {
		t.Errorf("Expected authenticated to be false, got %v", s.Values["authenticated"])
	}
	if s.Values["user_id"] != nil {
		t.Errorf("Expected user_id to be <nil>, got %v", s.Values["user_id"])
	}
}
