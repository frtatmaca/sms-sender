package redis_client

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
)

type client struct {
	Client *redis.Client
}

type Client interface {
	GracefulShutdown(graceTime time.Duration)
	GetRedisClient() *redis.Client
}

func NewClient(addr string) (Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return &client{Client: rdb}, nil
}

func (client *client) GracefulShutdown(graceTime time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), graceTime)
	defer func() {
		cancel()
		_ = client.Client.Close()
	}()
}

func (client *client) GetRedisClient() *redis.Client {
	return client.Client
}
