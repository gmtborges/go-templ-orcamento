package handlers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"

	"github.com/gustavomtborges/orcamento-auto/models"
	"github.com/gustavomtborges/orcamento-auto/services"
)

type MockUserStore struct {
	MockFn func() (*models.User, error)
}

func (m *MockUserStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return m.MockFn()
}

func TestLoginHandler_Create_Success(t *testing.T) {
	e := echo.New()
	authSvc := services.NewAuthService(&MockUserStore{})
	hash, _ := authSvc.GeneratePasswordHash("passwd123")
	st := &MockUserStore{
		MockFn: func() (*models.User, error) {
			return &models.User{Password: hash}, nil
		},
	}
	authSvc = services.NewAuthService(st)
	h := NewLoginHandler(authSvc)

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("email=test@example.com&password=passwd123"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("_session_store", sessions.NewCookieStore([]byte("secret")))

	if err := h.Create(c); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if rec.Code != http.StatusSeeOther {
		t.Errorf("Expected status See Other, got %v", rec.Code)
	}
}

func TestLoginHandler_Create_UserNotFound(t *testing.T) {
	e := echo.New()
	st := &MockUserStore{
		MockFn: func() (*models.User, error) {
			return nil, sql.ErrNoRows
		},
	}
	authSvc := services.NewAuthService(st)
	h := NewLoginHandler(authSvc)

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("email=test@example.com&password=passwd123"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.Create(c); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status Internal Server Error, got %v", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "E-mail e/ou senha incorretos") {
		t.Errorf("Expected body to contain E-mail e/ou senha incorretos, got %v", rec.Body.String())
	}
}

func TestLoginHandler_Create_InvalidPassword(t *testing.T) {
	e := echo.New()
	st := &MockUserStore{
		MockFn: func() (*models.User, error) {
			return &models.User{Password: "wrong"}, nil
		},
	}
	authSvc := services.NewAuthService(st)
	h := NewLoginHandler(authSvc)
	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		strings.NewReader("email=test@example.com&password=passwd123"),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.Create(c); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status Bad Request, got %v", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "E-mail e/ou senha incorretos") {
		t.Errorf("Expected body to contain E-mail e/ou senha incorretos, got %v", rec.Body.String())
	}
}

func TestLoginHandler_Create_Error_Database(t *testing.T) {
	e := echo.New()
	st := &MockUserStore{
		MockFn: func() (*models.User, error) {
			return nil, errors.New("error on database")
		},
	}
	authSvc := services.NewAuthService(st)
	h := NewLoginHandler(authSvc)
	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		strings.NewReader("email=test@example.com&password=passwd123"),
	)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h.Create(c); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("Expected status Internal Server Error, got %v", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Erro ao realizar login.") {
		t.Errorf("Expected body to contain Erro ao realizar login., got %v", rec.Body.String())
	}
}
