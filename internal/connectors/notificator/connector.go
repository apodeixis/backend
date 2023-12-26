package notificator

import (
	"context"
)

const createNotificationPath = "notifications"

func (c *connector) sendEmail(topic, address string, payload messageAttrs) error {
	if c.disabled {
		c.log.Warnf("disabled, skipping sending %s email to %s", topic, address)
		return nil
	}
	request := composeCreateEmailNotificationRequest(topic, address, payload)
	endpoint := c.endpoint.JoinPath(createNotificationPath)
	return c.connector.PostJSON(endpoint, request, context.Background(), nil)
}

func composeCreateEmailNotificationRequest(topic string, address string, payload messageAttrs) *createNotificationRequest {
	const (
		createNotificationType          = "create_notification"
		channelEmail                    = "email"
		notificationDestinationEmail    = "destination_type_email"
		typeNotificationMessageTemplate = "message_type_template"
	)
	channel := channelEmail
	request := &createNotificationRequest{
		Data: createNotificationData{
			Type: createNotificationType,
			Attributes: createNotificationAttributes{
				Channel: &channel,
				Message: message{
					Type:       typeNotificationMessageTemplate,
					Attributes: payload,
				},
				Topic: topic,
			},
			Relationships: createNotificationRelationships{
				Destinations: createNotificationDestinations{
					Data: []destinationKey{
						{
							Id:   address,
							Type: notificationDestinationEmail,
						},
					},
				},
			},
		},
	}
	return request
}
