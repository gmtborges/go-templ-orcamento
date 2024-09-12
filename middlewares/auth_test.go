package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/gmtborges/orcamento-auto/repo"
	"github.com/gmtborges/orcamento-auto/svc"
	"github.com/gmtborges/orcamento-auto/types"
)

func TestAuthtentication_NoSession(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	store := sessions.NewCookieStore([]byte("secret"))
	userRepo := &repo.MockUserRepository{
		MockFn: func() (interface{}, error) {
			return types.UserAuth{}, nil
		},
	}
	loginHandler := func(c echo.Context) error {
		return c.String(http.StatusOK, "session")
	}
	mw := session.Middleware(store)
	userSvc := svc.NewUserService(userRepo)
	protectedHandler := mw(Authentication(userSvc)(loginHandler))

	err := protectedHandler(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, rec.Code)
	assert.Equal(t, "/login", rec.Header().Get("Location"))
}

func TestAuthtentication_NoUserID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	store := sessions.NewCookieStore([]byte("secret"))
	userRepo := &repo.MockUserRepository{
		MockFn: func() (interface{}, error) {
			return types.UserAuth{}, nil
		},
	}
	loginHandler := func(c echo.Context) error {
		sess, _ := session.Get("auth-session", c)
		sess.Values["user_id"] = ""
		sess.Save(c.Request(), c.Response())
		return c.String(http.StatusOK, "session")
	}
	mw := session.Middleware(store)
	userSvc := svc.NewUserService(userRepo)
	protectedHandler := mw(Authentication(userSvc)(loginHandler))

	err := protectedHandler(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusSeeOther, rec.Code)
	assert.Equal(t, "/login", rec.Header().Get("Location"))
}
