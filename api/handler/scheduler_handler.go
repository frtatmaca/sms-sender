package handler

import (
	"net/http"

	"github.com/frtatmaca/sms-sender/api/scheduler"
	"github.com/gin-gonic/gin"
)

type SchedulerHandler struct {
	Scheduler *scheduler.Scheduler
}

func NewSchedulerHandler(scheduler *scheduler.Scheduler) *SchedulerHandler {
	return &SchedulerHandler{Scheduler: scheduler}
}

// SchedulerStart godoc
// @Summary Scheduler Start
// @Description Scheduler Start
// @Tags Sms
// @Produce json
// @Success 200
// @Router /api/v1/cronjob/start [get]
func (c *SchedulerHandler) SchedulerStart(ctx *gin.Context) {
	c.Scheduler.Start()
	ctx.Status(http.StatusOK)
}

// SchedulerStop godoc
// @Summary Scheduler Stop
// @Description Scheduler Stop
// @Tags Sms
// @Produce json
// @Success 200
// @Router /api/v1/cronjob/stop [get]
func (c *SchedulerHandler) SchedulerStop(ctx *gin.Context) {
	c.Scheduler.Stop()
	ctx.Status(http.StatusOK)
}
