package services

import (
	"context"

	"github.com/gustavomtborges/orcamento-auto/models"
	"github.com/gustavomtborges/orcamento-auto/repositories"
)

type LoginService struct {
	userRepo repositories.UserRepository
}

func NewLoginService(userRepo repositories.UserRepository) *LoginService {
	return &LoginService{userRepo: userRepo}
}

func (s *LoginService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
