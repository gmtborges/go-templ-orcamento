package services

import (
	"context"

	"github.com/gustavomtborges/orcamento-auto/models"
	"github.com/gustavomtborges/orcamento-auto/stores"
)

type LoginService struct {
	userStore stores.UserStorer
}

func NewLoginService(userStore stores.UserStorer) *LoginService {
	return &LoginService{userStore: userStore}
}

func (s *LoginService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.userStore.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
