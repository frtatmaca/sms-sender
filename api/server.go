package api

import (
	"github.com/frtatmaca/sms-sender/api/handler"
	"github.com/frtatmaca/sms-sender/api/middleware"
	"github.com/frtatmaca/sms-sender/api/scheduler"
	"github.com/frtatmaca/sms-sender/api/service"
	"github.com/frtatmaca/sms-sender/api/storage"
	_ "github.com/frtatmaca/sms-sender/swagger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func NewSmsSender(smsStorage storage.IStorage, scheduler *scheduler.Scheduler, logger *zap.SugaredLogger) *SmsSenderAPI {
	return &SmsSenderAPI{SmsStorage: smsStorage, Scheduler: scheduler, Logger: logger}
}

type SmsSenderAPI struct {
	SmsStorage       storage.IStorage
	SmsHandler       *handler.SmsHandler
	SchedulerHandler *handler.SchedulerHandler
	Scheduler        *scheduler.Scheduler
	Logger           *zap.SugaredLogger
}

func (n *SmsSenderAPI) Configure(router *gin.Engine) *gin.Engine {
	n.configureServices()
	n.configureRoutes(router)

	return router
}

func (n *SmsSenderAPI) configureServices() {
	// service
	smsService := service.NewSmsService(n.SmsStorage, n.Logger)

	// Controllers
	n.SmsHandler = handler.NewSmsHandler(smsService)
	n.SchedulerHandler = handler.NewSchedulerHandler(n.Scheduler)
}

func (n *SmsSenderAPI) configureRoutes(router *gin.Engine) {
	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/api/v1/notifications/sms", n.SmsHandler.ListSms)
	router.POST("/api/v1/notifications/sms", middleware.ValidateSmsRequest(), n.SmsHandler.SendSmsV1)
}
