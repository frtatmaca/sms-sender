package middleware

import (
	"net/http"

	"github.com/frtatmaca/sms-sender/api/model/request"

	"github.com/gin-gonic/gin"
)

func ValidateSmsRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var smsRequest request.SmsRequestV1
		if err := c.ShouldBindJSON(&smsRequest); err == nil {
			if len(smsRequest.Content) > 50 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Content length exceeds 50 characters"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
