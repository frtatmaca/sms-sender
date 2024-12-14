package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/frtatmaca/sms-sender/api/domain/entity"
	redis_client "github.com/frtatmaca/sms-sender/pkg/storage/redis"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type IStorage interface {
	List(ctx context.Context) ([]entity.Sms, error)
	Get(ctx context.Context, Id string) (*entity.Sms, error)
	Create(ctx context.Context, entity entity.Sms) error
	Update(ctx context.Context, entity *entity.Sms) error
	ReleaseLock(ctx context.Context, sms *entity.Sms) error
	GetAll(ctx context.Context) ([]entity.Sms, error)
}

type smsStorage struct {
	client *redis.Client
	logger *zap.SugaredLogger
}

func NewStorage(redisClient redis_client.Client, logger *zap.SugaredLogger) IStorage {
	return &smsStorage{client: redisClient.GetRedisClient(), logger: logger}
}

func (storage *smsStorage) Create(ctx context.Context, sms entity.Sms) error {
	data, err := json.Marshal(sms)
	if err != nil {
		panic(err)
	}

	if _, err := storage.client.HSet(ctx, "activeStatus:true", sms.Id.String(), data).Result(); err != nil {
		storage.logger.Errorw(fmt.Sprintf("Error while sms create %s", sms.Id), zap.Any("error", err))
		return err
	}
	return nil
}

func (storage *smsStorage) List(ctx context.Context) ([]entity.Sms, error) {
	result, err := storage.client.HGetAll(ctx, "activeStatus:true").Result()
	if err != nil {
		return nil, err
	}

	count := 0
	documents := make([]entity.Sms, 0)
	for _, value := range result {
		var msg entity.Sms
		_ = json.Unmarshal([]byte(value), &msg)

		if msg.ActiveStatus == true {
			lockKey := msg.Id.String() + ":lock"
			ttl := 10 * time.Second // 5 saniyelik TTL

			success, err := storage.client.SetNX(ctx, lockKey, "locked", ttl).Result()
			if err != nil {
				storage.logger.Errorw(fmt.Sprintf("Error setting locke %s", msg.Id), zap.Any("error", err))
				continue
			}

			if success {
				storage.logger.Errorw(fmt.Sprintf("Lock set for key %s", msg.Id), zap.Any("error", err))
				count++
			}
		}

		documents = append(documents, msg)

		if count == 100 {
			break
		}
	}

	return documents, nil
}

func (storage *smsStorage) Get(ctx context.Context, Id string) (*entity.Sms, error) {
	val, err := storage.client.HGet(ctx, "activeStatus:true", Id).Result()
	if err != nil || err == redis.Nil {
		return nil, nil
	}

	var msg *entity.Sms
	_ = json.Unmarshal([]byte(val), &msg)

	return msg, nil
}

func (storage *smsStorage) Update(ctx context.Context, sms *entity.Sms) error {
	data, _ := json.Marshal(sms)

	if _, err := storage.client.HSet(ctx, "activeStatus:false", sms.Id.String(), data).Result(); err != nil {
		storage.logger.Errorw(fmt.Sprintf("Error while add active status false %s", sms.Id), zap.Any("error", err))
		return err
	}

	if _, err := storage.client.HDel(ctx, "activeStatus:true", sms.Id.String()).Result(); err != nil {
		storage.logger.Errorw(fmt.Sprintf("Error while delete active status true %s", sms.Id), zap.Any("error", err))
		return err
	}

	return nil
}

func (storage *smsStorage) ReleaseLock(ctx context.Context, sms *entity.Sms) error {
	lockKey := sms.Id.String() + ":lock"
	_, err := storage.client.Del(ctx, lockKey).Result()
	return err
}

func (storage *smsStorage) GetAll(ctx context.Context) ([]entity.Sms, error) {
	documents := make([]entity.Sms, 0)
	keys := []string{"activeStatus:false"}

	for _, key := range keys {
		result, err := storage.client.HGetAll(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		for _, value := range result {
			var msg entity.Sms
			_ = json.Unmarshal([]byte(value), &msg)

			documents = append(documents, msg)
		}
	}

	return documents, nil
}
