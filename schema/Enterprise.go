package schema

type Enterprise struct {
	Name          string `json: "name"`
	PhoneNumberId string `json: "phoneNumberId"`
	Token         string `json: "token"`
}
