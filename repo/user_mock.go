package repo

import (
	"context"
	"fmt"

	"github.com/gmtborges/orcamento-auto/types"
)

type MockUserRepository struct {
	MockFn func() (interface{}, error)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*types.UserAuth, error) {
	data, err := m.MockFn()
	if data == nil {
		return nil, err
	}
	userAuth, ok := data.(types.UserAuth)
	if !ok {
		return nil, fmt.Errorf("failed to assert data as types.UserAuth")
	}
	return &userAuth, err
}

func (m *MockUserRepository) GetByID(ctx context.Context, userID int64) (int64, error) {
	data, err := m.MockFn()
	return data.(int64), err
}
