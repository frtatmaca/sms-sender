package service_test

import (
	"context"
	"github.com/frtatmaca/sms-sender/api/domain/entity"
	"github.com/frtatmaca/sms-sender/api/service"
	"github.com/frtatmaca/sms-sender/api/storage/mocks"
	client_mocks "github.com/frtatmaca/sms-sender/pkg/sms_sender/mocks"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"
)

func Test_Send(t *testing.T) {
	t.Parallel()

	t.Run("it returns err nil Send", func(t *testing.T) {
		mockStorage := new(mocks.MockStorage)
		mockClient := new(client_mocks.MockClient)
		logger := zap.NewNop().Sugar()
		senderService := service.NewSmsSenderService(mockStorage, mockClient, logger)

		ctx := context.Background()
		smsList := []entity.Sms{
			{To: "1234567890", Content: "Test message 1"},
			{To: "0987654321", Content: "Test message 2"},
		}

		mockStorage.On("List", ctx).Return(smsList, nil)
		mockClient.On("SmsSend", "1234567890", "Test message 1").Return("msgid1", nil)
		mockClient.On("SmsSend", "0987654321", "Test message 2").Return("msgid2", nil)
		mockStorage.On("ReleaseLock", ctx, mock.AnythingOfType("*entity.Sms")).Return(nil)
		mockStorage.On("Update", ctx, mock.AnythingOfType("*entity.Sms")).Return(nil)

		senderService.Send(ctx)

		mockStorage.AssertExpectations(t)
		mockClient.AssertExpectations(t)
	})
}
