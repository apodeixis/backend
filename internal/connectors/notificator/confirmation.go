package notificator

type ConfirmationPayload struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Link string `json:"link"`
}

func (c *connector) SendConfirmationEmail(address string, payload *ConfirmationPayload) error {
	const topic = "confirm"
	return c.sendEmail(topic, address, messageAttrs{Payload: payload})
}
