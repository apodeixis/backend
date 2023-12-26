package responses

import (
	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/resources"
)

type UserWithAuthTokenResponse struct {
	Data     resources.User                          `json:"data"`
	Included resources.Login200ResponseIncludedInner `json:"included"`
}

func ComposeUserWithAuthToken(user *data.User, authToken *resources.AuthToken) *UserWithAuthTokenResponse {
	userResource := *convertToUserResource(user)
	userResource.Relationships = &resources.UserAllOfRelationships{
		Tokens: resources.NewUserAllOfRelationshipsTokens(resources.AuthTokenKey{Type: "auth_token"}),
	}
	return &UserWithAuthTokenResponse{
		Data: userResource,
		Included: resources.Login200ResponseIncludedInner{
			AuthToken: authToken,
		},
	}
}
