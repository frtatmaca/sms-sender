package handler

import (
	"net/http"

	"github.com/frtatmaca/sms-sender/api/error/httperror"
	"github.com/frtatmaca/sms-sender/api/model/request"
	"github.com/frtatmaca/sms-sender/api/service"
	"github.com/gin-gonic/gin"
)

type SmsHandler struct {
	SmsService service.ISmsService
}

func NewSmsHandler(smsService service.ISmsService) *SmsHandler {
	return &SmsHandler{SmsService: smsService}
}

// SendSmsV1 godoc
//
//	@Summary		send sms
//	@Description	send sms
//	@Tags			SMS
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.SmsRequestV1				true	"SMS you want to create"
//	@Success		201		{object}	entity.Sms					"The newly created SMS"
//	@Router			/api/v1/notifications/sms [post]
func (h *SmsHandler) SendSmsV1(ctx *gin.Context) {
	var req request.SmsRequestV1
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(httperror.NewWithMetadata(httperror.InvalidRequest, err.Error()))
		return
	}

	sms, err := h.SmsService.Create(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSONP(http.StatusCreated, &sms)
}

// ListSms godoc
// @Summary List Sms
// @Description List Sms
// @Tags Sms
// @Produce json
// @Success 200 {array} entity.Sms
// @Router /api/v1/notifications/sms [get]
func (c *SmsHandler) ListSms(ctx *gin.Context) {
	data, err := c.SmsService.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}
