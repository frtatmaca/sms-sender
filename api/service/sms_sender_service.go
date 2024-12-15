package service

import (
	"context"
	"fmt"
	"github.com/frtatmaca/sms-sender/api/domain/entity"
	"github.com/frtatmaca/sms-sender/api/storage"
	sms_sender "github.com/frtatmaca/sms-sender/pkg/sms_sender/telefonica"
	"go.uber.org/zap"
	"sync"
)

var _ ISmsSenderService = new(SenderService)

//go:generate mockery --name=ISmsSenderService
type ISmsSenderService interface {
	Send(ctx context.Context)
}

type SenderService struct {
	SmsStorage      storage.IStorage
	SmsSenderClient sms_sender.IClient
	logger          *zap.SugaredLogger
}

func NewSmsSenderService(smsStorage storage.IStorage, smsSenderClient sms_sender.IClient, logger *zap.SugaredLogger) *SenderService {
	return &SenderService{SmsStorage: smsStorage, SmsSenderClient: smsSenderClient, logger: logger}
}

func (s *SenderService) Send(ctx context.Context) {
	list, err := s.SmsStorage.List(ctx)
	if err != nil {
		s.logger.Errorw(fmt.Sprintf("Error while sms fetching: "), zap.Any("error", err))
	}
	chunkSize := 10
	var wg sync.WaitGroup

	for i := 0; i < len(list); i += chunkSize {
		end := i + chunkSize
		if end > len(list) {
			end = len(list)
		}
		chunk := list[i:end]
		wg.Add(1)
		go processChunk(ctx, chunk, &wg, s.SmsStorage, s.SmsSenderClient, s.logger)
	}

	wg.Wait()
}

func processChunk(ctx context.Context, chunk []entity.Sms, wg *sync.WaitGroup, smsStorage storage.IStorage, smsSenderClient sms_sender.IClient, logger *zap.SugaredLogger) {
	defer wg.Done()
	for _, num := range chunk {
		messageId, err := smsSenderClient.SmsSend(num.To, num.Content)
		if err != nil {
			logger.Errorw(fmt.Sprintf("Error while sms sending %s", num.Id), zap.Any("error", err))
			return
		}

		err = smsStorage.ReleaseLock(ctx, &num)
		if err != nil {
			logger.Errorw(fmt.Sprintf("Error while release lock %s", num.Id), zap.Any("error", err))
		}

		num.DeActivate(messageId)
		
		err = smsStorage.Update(ctx, &num)
		if err != nil {
			logger.Errorw(fmt.Sprintf("Error while record updateting %s", num.Id), zap.Any("error", err))
		}
	}
}
