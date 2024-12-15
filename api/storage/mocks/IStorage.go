package mocks

import (
	"context"

	"github.com/frtatmaca/sms-sender/api/domain/entity"
	"github.com/stretchr/testify/mock"
)

// MockStorage is a mock implementation of the IStorage interface
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) List(ctx context.Context) ([]entity.Sms, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Sms), args.Error(1)
}

func (m *MockStorage) Get(ctx context.Context, Id string) (*entity.Sms, error) {
	args := m.Called(ctx, Id)
	return args.Get(0).(*entity.Sms), args.Error(1)
}

func (m *MockStorage) Create(ctx context.Context, sms entity.Sms) error {
	args := m.Called(ctx, sms)
	return args.Error(0)
}

func (m *MockStorage) Update(ctx context.Context, sms *entity.Sms) error {
	args := m.Called(ctx, sms)
	return args.Error(0)
}

func (m *MockStorage) ReleaseLock(ctx context.Context, sms *entity.Sms) error {
	args := m.Called(ctx, sms)
	return args.Error(0)
}

func (m *MockStorage) GetAll(ctx context.Context) ([]entity.Sms, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Sms), args.Error(1)
}
