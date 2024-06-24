package whatsapp

func Send(recipient, body string) {
	Params.SetTo(recipient)
	Params.SetBody(body)

}
