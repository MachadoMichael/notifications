package schema

type Enterprise struct {
	Name          string `json: "name"`
	PhoneNumberId string `json: "phoneNumberId"`
	Token         string `json: "token"`
	Email         string `json: "email"`
	Conversations int    `json: "conversations"`
	Messages      int    `json: "messages"`
}
