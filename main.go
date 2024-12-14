package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/frtatmaca/sms-sender/api"
	"github.com/frtatmaca/sms-sender/api/config"
	"github.com/frtatmaca/sms-sender/api/scheduler"
	"github.com/frtatmaca/sms-sender/api/storage"
	"github.com/frtatmaca/sms-sender/pkg/logging"
	sms_sender "github.com/frtatmaca/sms-sender/pkg/sms_sender/telefonica"
	"github.com/frtatmaca/sms-sender/pkg/storage/redis"
	"github.com/gin-gonic/gin"
)

const (
	shutdownTimeout = 10 * time.Second
)

func main() {
	cfg := config.NewConfiguration()

	logger := logging.NewLoggerWithLevel("stderr", cfg.LogLevel)
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Fatalf("error syncing logs:  %v", err)
		}
	}()

	fmt.Print(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	routeHandler := gin.New()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// client
	smsSenderClint, _ := sms_sender.NewClient()
	redisClient, err := redis_client.NewClient(cfg.Redis.Addr)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer redisClient.GracefulShutdown(10 * time.Second)

	smsStorage := storage.NewStorage(redisClient, logger)

	// Scheduler
	newScheduler := scheduler.NewScheduler(cfg, smsStorage, logger, smsSenderClint)
	if err != nil {
		logger.Fatal(err.Error())
	}
	newScheduler.Run(ctx)

	smsSender := api.NewSmsSender(smsStorage, newScheduler, logger)
	smsSender.Configure(routeHandler)

	srv := &http.Server{
		Addr:    ":8050",
		Handler: routeHandler,

		ReadHeaderTimeout: 20 * time.Second,
	}

	go func() {
		logger.Infof("Starting HTTP server on: %s", "localhost:8050")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed to start server")
		}
	}()

	<-ctx.Done()

	stop()
	logger.Infof("Interrupt signal received, initiating shutdown with process exit in: %v", shutdownTimeout)

	timeoutContext, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(timeoutContext); err != nil {
		logger.Error("server forced to shutdown")
	}
}
