package services

import (
	"context"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/repos"
	"github.com/gmtborges/orcamento-auto/types"
)

type UserService struct {
	userRepo repos.UserRepository
}

func NewUserService(userRepo repos.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*types.UserAuth, error) {
	return s.userRepo.GetUserByEmail(ctx, email)
}

func (s *UserService) GetByID(ctx context.Context, uID int64) (int64, error) {
	return s.userRepo.GetByID(ctx, uID)
}

func (s *UserService) SetSession(c echo.Context, companyID int64, userID int64, roles []string) error {
	session, err := session.Get("session", c)
	if err != nil {
		return err
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	session.Values["userID"] = userID
	session.Values["companyID"] = companyID
	session.Values["roles"] = roles
	return session.Save(c.Request(), c.Response())
}

func (s *UserService) IsAuthenticated(c echo.Context) (bool, error) {
	session, err := session.Get("session", c)
	if err != nil {
		return false, err
	}
	uID, ok := session.Values["userID"]
	return uID != nil && ok, nil
}

func (s *UserService) RemoveSession(c echo.Context) error {
	session, err := session.Get("session", c)
	if err != nil {
		return err
	}
	session.Values["userID"] = nil
	session.Values["companyID"] = nil
	session.Values["roles"] = nil
	session.Options.MaxAge = -1
	return session.Save(c.Request(), c.Response())
}
