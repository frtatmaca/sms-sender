package scheduler

import (
	"context"
	"time"

	"github.com/frtatmaca/sms-sender/api/config"
	"github.com/frtatmaca/sms-sender/api/service"
	"github.com/frtatmaca/sms-sender/api/storage"
	sms_sender "github.com/frtatmaca/sms-sender/pkg/sms_sender/telefonica"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Scheduler struct {
	cfg             *config.AppConfig
	smsStorage      storage.IStorage
	cron            *cron.Cron
	logger          *zap.SugaredLogger
	SmsSenderClient sms_sender.IClient
	jobs            []Job
}

func NewScheduler(cfg *config.AppConfig, smsStorage storage.IStorage, logger *zap.SugaredLogger, smsSenderClient sms_sender.IClient) *Scheduler {
	jobs := initializeJobs(cfg, smsStorage, logger, smsSenderClient)
	return &Scheduler{
		cfg:             cfg,
		cron:            cron.New(),
		logger:          logger,
		smsStorage:      smsStorage,
		SmsSenderClient: smsSenderClient,
		jobs:            jobs,
	}
}

func initializeJobs(cfg *config.AppConfig, smsStorage storage.IStorage, logger *zap.SugaredLogger, smsSenderClient sms_sender.IClient) []Job {
	var jobs []Job
	senderService := service.NewSmsSenderService(smsStorage, smsSenderClient, logger)
	jobs = append(jobs, Job{SenderService: senderService, cronExpression: cfg.CronExpression, name: "SmsSender"})

	return jobs
}

func (s *Scheduler) Run(ctx context.Context) {
	for _, job := range s.jobs {
		_, err := s.cron.AddFunc(job.cronExpression, func() {
			s.logger.Infof("Shovel job is running for %s", job.name)
			job.Send(ctx)
		})
		if err != nil {
			s.logger.Fatalf("Error during cronjob creation %+v", err)
		}
		s.cron.Start()
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}

type Job struct {
	*service.SenderService
	name           string
	runningTime    time.Duration
	cronExpression string
}
