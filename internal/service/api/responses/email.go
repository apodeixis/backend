package responses

import "github.com/apodeixis/backend/resources"

type EmailResponse struct {
	Data resources.Email `json:"data"`
}

func ComposeEmail(email string) *EmailResponse {
	return &EmailResponse{
		Data: resources.Email{
			Attributes: resources.EmailAllOfAttributes{
				Email: email,
			},
		},
	}
}
