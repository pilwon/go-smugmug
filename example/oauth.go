package main

import (
	"net/http"

	"github.com/mrjones/oauth"
)

const (
	RequestTokenUrl   = "https://api.smugmug.com/services/oauth/1.0a/getRequestToken"
	AuthorizeTokenUrl = "https://api.smugmug.com/services/oauth/1.0a/authorize"
	AccessTokenUrl    = "https://api.smugmug.com/services/oauth/1.0a/getAccessToken"

	URLParamAccessFull        = "Full"
	URLParamPermissionsModify = "Modify"
)

func buildOAuthHTTPClient(consumerKey, consumerSecret, accessToken, accessTokenSecret string) (*http.Client, error) {
	consumer := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   RequestTokenUrl,
			AuthorizeTokenUrl: AuthorizeTokenUrl,
			AccessTokenUrl:    AccessTokenUrl,
		},
	)
	consumer.AdditionalAuthorizationUrlParams = map[string]string{
		"Access":      URLParamAccessFull,
		"Permissions": URLParamPermissionsModify,
	}

	token := &oauth.AccessToken{
		Token:  accessToken,
		Secret: accessTokenSecret,
	}

	client, err := consumer.MakeHttpClient(token)
	if err != nil {
		return nil, err
	}

	return client, nil
}
