package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


func SetUpConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID: "707879399164-esp7a23mv1dnkl6d7asaj6jpdbhtjf37.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL: "http://localhost:8081/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}