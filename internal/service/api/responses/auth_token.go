package responses

import "github.com/apodeixis/backend/resources"

type AuthTokenResponse struct {
	Data resources.AuthToken `json:"data"`
}

func ComposeAuthToken(authToken *resources.AuthToken) *AuthTokenResponse {
	return &AuthTokenResponse{
		Data: *authToken,
	}
}
