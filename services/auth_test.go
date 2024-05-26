package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func TestAuthService_Generate_Verify_Password(t *testing.T) {
	authSvc := NewAuthService()
	hash, err := authSvc.GeneratePasswordHash("passwd123")
	if err != nil {
		t.Errorf("Error generating hash: %v", err)
	}
	isValid, err := authSvc.VerifyPasswordHash("passwd123", hash)
	if err != nil {
		t.Errorf("Error verifying password: %v", err)
	}
	if !isValid {
		t.Errorf("Expected isValid to be true, got %v", isValid)
	}
}

type MockSessionStore struct {
	session *sessions.Session
	err     error
}

func (m *MockSessionStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	return m.session, m.err
}

func TestAuthService_SetAuthSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	svc := NewAuthService()
	err := svc.SetAuthSession(c, 123)
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
	if s.Values["user_id"] != 123 {
		t.Errorf("Expected user_id to be 123, got %v", s.Values["user_id"])
	}
}

func TestAuthService_RemoveAuthSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	svc := NewAuthService()
	if err := svc.SetAuthSession(c, 123); err != nil {
		t.Errorf("Error setting session %v", err)
	}
	err := svc.RemoveAuthSession(c)
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
