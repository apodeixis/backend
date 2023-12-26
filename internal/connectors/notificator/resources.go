package notificator

type (
	createNotificationRequest struct {
		Data createNotificationData `json:"data"`
	}
	createNotificationData struct {
		Type          string                          `json:"type"`
		Attributes    createNotificationAttributes    `json:"attributes"`
		Relationships createNotificationRelationships `json:"relationships"`
	}
	createNotificationAttributes struct {
		Channel *string `json:"channel,omitempty"`
		Message message `json:"message"`
		Topic   string  `json:"topic"`
	}
	message struct {
		Type       string       `json:"type"`
		Attributes messageAttrs `json:"attributes"`
	}
	messageAttrs struct {
		Payload interface{} `json:"payload"`
	}
	createNotificationRelationships struct {
		Destinations createNotificationDestinations `json:"destinations"`
	}
	createNotificationDestinations struct {
		Data []destinationKey `json:"data"`
	}
	destinationKey struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	}
)
