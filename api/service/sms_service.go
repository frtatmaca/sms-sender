package service

import (
	"context"
	"fmt"

	"github.com/frtatmaca/sms-sender/api/domain/entity"
	"github.com/frtatmaca/sms-sender/api/model/request"
	"github.com/frtatmaca/sms-sender/api/storage"
	"go.uber.org/zap"
)

var _ ISmsService = new(Service)

//go:generate mockery --name=ISmsService
type ISmsService interface {
	Create(ctx context.Context, input *request.SmsRequestV1) (*entity.Sms, error)
	GetAll(ctx context.Context) ([]entity.Sms, error)
}

type Service struct {
	SmsStorage storage.IStorage
	logger     *zap.SugaredLogger
}

func NewSmsService(smsStorage storage.IStorage, logger *zap.SugaredLogger) *Service {
	return &Service{SmsStorage: smsStorage, logger: logger}
}

func (s *Service) Create(ctx context.Context, input *request.SmsRequestV1) (*entity.Sms, error) {
	smsEntity := entity.NewSms(input.To, input.Content)
	err := s.SmsStorage.Create(ctx, smsEntity)
	if err != nil {
		s.logger.Errorw(fmt.Sprintf("Error while sms creating: "), zap.Any("error", err))

		return nil, err
	}

	return &smsEntity, nil
}

func (s *Service) GetAll(ctx context.Context) ([]entity.Sms, error) {
	list, err := s.SmsStorage.GetAll(ctx)
	if err != nil {
		s.logger.Errorw(fmt.Sprintf("Error while sms fetching: "), zap.Any("error", err))

		return nil, err
	}

	return list, nil
}
