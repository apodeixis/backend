package helpers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/apodeixis/backend/internal/connectors/notificator"
	"github.com/apodeixis/backend/internal/service/api/ctx"
)

func ComposeConfirmationPayload(name, token string, expiresAt time.Time, r *http.Request) *notificator.ConfirmationPayload {
	values := ctx.WebConfig(r).ConfirmationEndpoint.Query()
	const queryTokenKey = "token"
	values.Set(queryTokenKey, token)
	values.Encode()
	ctx.WebConfig(r).ConfirmationEndpoint.RawQuery = values.Encode()
	lifeMinutes := expiresAt.Add(-time.Duration(time.Now().UTC().Minute() * int(time.Minute))).Minute()
	return &notificator.ConfirmationPayload{
		Name: name,
		Time: fmt.Sprintf("%d minute(s)", lifeMinutes),
		Link: ctx.WebConfig(r).ConfirmationEndpoint.String(),
	}
}

func ComposeRecoveryPayload(name, token string, expiresAt time.Time, r *http.Request) *notificator.RecoveryPayload {
	values := ctx.WebConfig(r).RecoveryEndpoint.Query()
	const queryTokenKey = "token"
	values.Set(queryTokenKey, token)
	values.Encode()
	ctx.WebConfig(r).RecoveryEndpoint.RawQuery = values.Encode()
	lifeMinutes := expiresAt.Add(-time.Duration(time.Now().UTC().Minute() * int(time.Minute))).Minute()
	return &notificator.RecoveryPayload{
		Name: name,
		Time: fmt.Sprintf("%d minute(s)", lifeMinutes),
		Link: ctx.WebConfig(r).RecoveryEndpoint.String(),
	}
}
