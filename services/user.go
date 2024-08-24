package services

import (
	"context"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/models"
	"github.com/gmtborges/orcamento-auto/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetRolesByUserID(userID int) ([]string, error) {
	return []string{"admin", "coop_admin"}, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) SetUserSession(c echo.Context, userID int64, roles string) error {
	session, err := session.Get("auth-session", c)
	if err != nil {
		return err
	}

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	session.Values["authenticated"] = true
	session.Values["user_id"] = userID
	session.Values["roles"] = roles
	return session.Save(c.Request(), c.Response())
}

func (s *UserService) IsAuthenticated(c echo.Context) (bool, error) {
	session, err := session.Get("auth-session", c)
	if err != nil {
		return false, err
	}
	auth, ok := session.Values["authenticated"].(bool)
	return auth && ok, nil
}

func (s *UserService) RemoveUserSession(c echo.Context) error {
	session, err := session.Get("auth-session", c)
	if err != nil {
		return err
	}
	session.Values["authenticated"] = false
	session.Values["user_id"] = nil
	session.Options.MaxAge = -1
	return session.Save(c.Request(), c.Response())
}
