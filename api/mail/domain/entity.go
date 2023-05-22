package domain

type SendEmailPayload struct {
	Recipient []string `json:"recipient"`
	Subject   string   `json:"subject"   validate:"required"`
	Body      string   `json:"body"      validate:"required"`
	SentBy    string   `json:"-"`
}
