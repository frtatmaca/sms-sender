package mocks

import "github.com/stretchr/testify/mock"

// MockClient is a mock implementation of the IClient interface
type MockClient struct {
	mock.Mock
}

func (m *MockClient) SmsSend(to string, content string) (string, error) {
	args := m.Called(to, content)
	return args.String(0), args.Error(1)
}
