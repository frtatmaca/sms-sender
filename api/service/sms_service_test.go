package service_test

import (
	"context"
	"errors"
	"github.com/frtatmaca/sms-sender/api/domain/entity"
	"github.com/frtatmaca/sms-sender/api/model/request"
	"github.com/frtatmaca/sms-sender/api/service"
	"github.com/frtatmaca/sms-sender/api/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"
)

func Test_Create(t *testing.T) {
	t.Parallel()

	t.Run("it returns err nil created sms", func(t *testing.T) {
		t.Parallel()
		mockStorage := new(mocks.MockStorage)
		logger := zap.NewNop().Sugar()
		smsService := service.NewSmsService(mockStorage, logger)

		ctx := context.Background()
		input := request.SmsRequestV1{
			To:      "1234567890",
			Content: "Test message",
		}
		smsEntity := entity.NewSms(input.To, input.Content)

		mockStorage.On("Create", ctx, mock.AnythingOfType("entity.Sms")).Return(nil)

		result, err := smsService.Create(ctx, &input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, smsEntity.To, result.To)
		assert.Equal(t, smsEntity.Content, result.Content)

		mockStorage.AssertExpectations(t)
	})

	t.Run("it returns err created sms", func(t *testing.T) {
		t.Parallel()
		mockStorage := new(mocks.MockStorage)
		logger := zap.NewNop().Sugar()
		smsService := service.NewSmsService(mockStorage, logger)

		ctx := context.Background()
		input := request.SmsRequestV1{
			To:      "1234567890",
			Content: "Test message",
		}

		mockStorage.On("Create", ctx, mock.AnythingOfType("entity.Sms")).Return(errors.New("some error"))

		_, err := smsService.Create(ctx, &input)

		assert.Error(t, err, errors.New("some error"))

		mockStorage.AssertExpectations(t)
	})
}

func Test_GetAll(t *testing.T) {
	t.Parallel()

	t.Run("it returns err nil Get All Sms", func(t *testing.T) {
		mockStorage := new(mocks.MockStorage)
		logger := zap.NewNop().Sugar()
		smsService := service.NewSmsService(mockStorage, logger)

		ctx := context.Background()
		expectedList := []entity.Sms{
			{To: "1234567890", Content: "Test message 1"},
			{To: "0987654321", Content: "Test message 2"},
		}

		mockStorage.On("GetAll", ctx).Return(expectedList, nil)

		result, err := smsService.GetAll(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedList, result)

		mockStorage.AssertExpectations(t)
	})

	t.Run("it returns err Get All Sms", func(t *testing.T) {
		mockStorage := new(mocks.MockStorage)
		logger := zap.NewNop().Sugar()
		smsService := service.NewSmsService(mockStorage, logger)

		ctx := context.Background()
		expectedList := []entity.Sms{
			{To: "1234567890", Content: "Test message 1"},
			{To: "0987654321", Content: "Test message 2"},
		}

		mockStorage.On("GetAll", ctx).Return(expectedList, errors.New("some error"))

		_, err := smsService.GetAll(ctx)

		assert.Error(t, err, errors.New("some error"))

		mockStorage.AssertExpectations(t)
	})
}
