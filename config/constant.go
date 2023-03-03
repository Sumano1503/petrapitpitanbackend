package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


func SetUpConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID: "490179235093-n8vurjhtci5qp13di65nd1g9qbqoevh7.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL: "http://localhost:8080/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}