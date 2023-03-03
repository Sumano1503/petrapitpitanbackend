package logincontroller

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request){
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

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

 http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request){

}