package schema

type MessageRequest struct {
	Recipient string `json: "recipient"`
	Body      string `json: "body"`
}
