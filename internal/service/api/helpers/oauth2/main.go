package oauth2

import (
	"context"

	"golang.org/x/oauth2"
)

func GetUserToken(code string, cfg *oauth2.Config) (*oauth2.Token, error) {
	return cfg.Exchange(context.Background(), code)
}
