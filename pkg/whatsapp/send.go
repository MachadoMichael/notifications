package whatsapp

import "os"

func Send(recipient, body string) {
	Params.SetTo(recipient)
	Params.SetBody(body)

	res, err := Client.Api.CreateMessage(Params)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	} else {
		if res.Body != nil {
			println(*res.Body)
		} else {
			println(res.Body)
		}
	}

}
