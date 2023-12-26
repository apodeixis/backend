package responses

import (
	"net/http"
	"strconv"

	"github.com/apodeixis/backend/internal/data"
	"github.com/apodeixis/backend/internal/service/api/ctx"
	"github.com/apodeixis/backend/internal/service/api/requests"
	"github.com/apodeixis/backend/resources"
)

func ComposePostHeaderListResponse(
	headers []data.PostHeader,
	r *http.Request,
) (*resources.GetConfirmedPostsHeaders200Response, error) {

	var (
		data          = make([]resources.PostHeader, len(headers))
		included      = make([]resources.User, len(headers))
		usersQ        = ctx.UsersQ(r)
		starredPostsQ = ctx.StarredPostsQ(r)
	)

	for i, header := range headers {
		user, err := usersQ.New().FilterByID(header.UserID).Get()
		if err != nil {
			return nil, err
		}

		userResource := *convertToUserResource(user)

		var (
			starred   bool
			claims, _ = requests.ExtractClaimsFromAuthHeader(r)
		)
		if claims != nil {
			starredForUserID := claims.OwnerId
			starredPost, err := starredPostsQ.New().FilterByPostID(header.ID).FilterByUserID(starredForUserID).Get()
			if err != nil {
				return nil, err
			}
			starred = starredPost != nil
		}

		headerResource := *convertPostHeaderToResource(&header, starred)
		headerResource.Relationships = &resources.PostAllOfRelationships{
			Author: resources.PostAllOfRelationshipsAuthor{
				Data: resources.UserKey{
					Id:   strconv.FormatInt(header.UserID, 10),
					Type: "user",
				},
			},
		}
		data[i] = headerResource
		included[i] = userResource
	}
	return &resources.GetConfirmedPostsHeaders200Response{
		Data:     data,
		Included: included,
	}, nil
}

func ComposeStarredPostHeaderListResponse(
	headers []data.PostHeader,
	usersQ data.Users,
) (*resources.GetConfirmedPostsHeaders200Response, error) {

	var (
		data     = make([]resources.PostHeader, len(headers))
		included = make([]resources.User, len(headers))
	)

	const starred = true

	for i, header := range headers {
		user, err := usersQ.New().FilterByID(header.UserID).Get()
		if err != nil {
			return nil, err
		}
		var (
			headerResource = *convertPostHeaderToResource(&header, starred)
			userResource   = *convertToUserResource(user)
		)
		headerResource.Relationships = &resources.PostAllOfRelationships{
			Author: resources.PostAllOfRelationshipsAuthor{
				Data: resources.UserKey{
					Id:   strconv.FormatInt(header.UserID, 10),
					Type: "user",
				},
			},
		}
		data[i] = headerResource
		included[i] = userResource
	}
	return &resources.GetConfirmedPostsHeaders200Response{
		Data:     data,
		Included: included,
	}, nil
}
