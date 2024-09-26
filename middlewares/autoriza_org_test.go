package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/repo"
	"github.com/gmtborges/orcamento-auto/svc"
)

func TestAuthorizeOrg(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/bidding", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userSvc := svc.NewUsuarioService(&repo.MockUsuarioRepository{})
	userSvc.SetSession(c, 321, 123, []string{"admin"})
}
