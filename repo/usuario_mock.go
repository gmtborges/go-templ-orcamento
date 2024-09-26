package repo

import (
	"context"
	"fmt"

	"github.com/gmtborges/orcamento-auto/types"
)

type MockUsuarioRepository struct {
	MockFn func() (interface{}, error)
}

func (m *MockUsuarioRepository) GetUsuarioByEmail(ctx context.Context, email string) (*types.UsuarioAutenticacao, error) {
	data, err := m.MockFn()
	if data == nil {
		return nil, err
	}
	userAuth, ok := data.(types.UsuarioAutenticacao)
	if !ok {
		return nil, fmt.Errorf("failed to assert data as types.UserAuth")
	}
	return &userAuth, err
}

func (m *MockUsuarioRepository) GetByID(ctx context.Context, userID int64) (int64, error) {
	data, err := m.MockFn()
	return data.(int64), err
}
