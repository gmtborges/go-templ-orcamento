package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/repositories"
	"github.com/gmtborges/orcamento-auto/services"
)

func TestHasRoleAllowAssoc(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/bidding", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userSvc := services.NewUserService(&repositories.MockUserRepository{})
	userSvc.SetSession(c, 123, "assoc")
}
