package sms_sender

type client struct{}

type IClient interface {
	SmsSend(to, message string) (string, error)
}

func NewClient() (IClient, error) {
	return &client{}, nil
}

func (client *client) SmsSend(to, message string) (string, error) {
	return "message-id", nil
}
