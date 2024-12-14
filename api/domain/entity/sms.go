package entity

import (
	"time"

	"github.com/google/uuid"
)

type Sms struct {
	Id           uuid.UUID `redis:"id" json:"id" example:"00000000-0000-0000-0000-000000000000"`
	To           string    `redis:"to" json:"to"`
	Content      string    `redis:"content" json:"content"`
	ActiveStatus bool      `redis:"activeStatus" json:"activeStatus"`
	MessageId    string    `redis:"messageId" json:"messageId"`
	CreatedAt    time.Time `redis:"createdAt" json:"createdAt"`
}

func NewSms(to string, content string) Sms {
	return Sms{
		Id:           uuid.New(),
		To:           to,
		Content:      content,
		ActiveStatus: true,
	}
}
