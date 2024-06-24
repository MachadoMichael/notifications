package whatsapp

func Sender(recipient, body string) error {
	Params.SetTo(recipient)
	Params.SetBody(body)

	res, err := Client.Api.CreateMessage(Params)
	if err != nil {
		return err
	} else {
		if res.Body != nil {
			println(*res.Body)
			return nil
		} else {
			println(res.Body)
			return nil
		}
	}

}
