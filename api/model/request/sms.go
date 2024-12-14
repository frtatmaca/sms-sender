package request

type SmsRequestV1 struct {
	To      string `json:"to"`
	Content string `json:"content"`
}
