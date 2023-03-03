package logincontroller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/config"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request){
 googleConfig := config.SetUpConfig()
 url := googleConfig.AuthCodeURL("randomstate")

 http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request){
	state := req.URL.Query () ["state" ] [0]
	if state != "randomstate" {
		fmt.Fprintln(res,"states dont match ")
		return
	}

	code := req.URL.Query()["code"][0]

	googleConfig := config.SetUpConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println(res, "CODE exchange failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println(res, "User Data Fetch Failed")
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(res, "JSON Parse Failed")
	}

	fmt.Println(res, string(userData))

}