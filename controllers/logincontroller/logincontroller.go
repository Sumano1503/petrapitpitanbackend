package logincontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/config"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request){
 googleConfig := config.SetUpConfig()
 url := googleConfig.AuthCodeURL("randomstate")

 http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request){

}