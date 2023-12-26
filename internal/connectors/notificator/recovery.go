package notificator

type RecoveryPayload struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Link string `json:"link"`
}

func (c *connector) SendRecoveryEmail(address string, payload *RecoveryPayload) error {
	const topic = "recovery"
	return c.sendEmail(topic, address, messageAttrs{Payload: payload})
}
