package notificator

import (
	"net/http"
	"net/url"

	jsonapi "gitlab.com/distributed_lab/json-api-connector"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/connectors/signed"
)

type Connector interface {
	SendRecoveryEmail(address string, payload *RecoveryPayload) error
	SendConfirmationEmail(address string, payload *ConfirmationPayload) error
}

type connector struct {
	disabled  bool
	log       *logan.Entry
	endpoint  *url.URL
	connector *jsonapi.Connector
}

func NewConnector(endpoint *url.URL, log *logan.Entry) Connector {
	client := signed.NewClient(http.DefaultClient, endpoint)
	return &connector{
		disabled:  false,
		log:       log.WithField("connector", "notificator"),
		endpoint:  endpoint,
		connector: jsonapi.NewConnector(client),
	}
}

func NewDisabledConnector(log *logan.Entry) Connector {
	return &connector{
		disabled: true,
		log:      log,
	}
}
