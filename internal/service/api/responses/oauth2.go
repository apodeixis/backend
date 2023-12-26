package responses

import "github.com/apodeixis/backend/resources"

func ComposeOAuth2Google(url string) *resources.OAuth2Google200Response {
	return &resources.OAuth2Google200Response{
		Data: resources.OAuth2{
			Type: "oauth2",
			Attributes: resources.OAuth2AllOfAttributes{
				Url: url,
			},
		},
	}
}
