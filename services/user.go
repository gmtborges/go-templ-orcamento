package services

import (
	"context"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/gmtborges/orcamento-auto/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

type UserAuth struct {
	ID       int64
	Name     string
	Password string
	roles    []string
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*UserAuth, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &UserAuth{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
		roles:    user.Roles,
	}, nil
}

func (s *UserService) GetByID(ctx context.Context, userID int64) (int64, error) {
	return s.userRepo.GetByID(ctx, userID)
}

func (s *UserService) SetSession(c echo.Context, userID int64, roles string) error {
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
