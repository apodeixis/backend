package responses

import "github.com/apodeixis/backend/resources"

func ComposePostsAmount(amount int64) *resources.GetPostsAmount200Response {
	return &resources.GetPostsAmount200Response{
		Data: resources.PostsAmount{
			Attributes: resources.PostsAmountAllOfAttributes{
				Amount: amount,
			},
		},
	}
}
