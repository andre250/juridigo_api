package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/juridigo/juridigo_api_usuario/helpers"
	"github.com/juridigo/juridigo_api_usuario/models"
	"github.com/juridigo/juridigo_api_usuario/utils"
)

/*
GetFacebookInfo - Responsável por retornar ifnormações do front
*/
func GetFacebookInfo(w http.ResponseWriter, r *http.Request) {
	if helpers.ReqRefuse(w, r, "POST") != nil {
		return
	}
	body, _ := facebookInfo(w, r)

	if body == nil {
		return
	}

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write(body)
}

func facebookInfo(w http.ResponseWriter, r *http.Request) ([]byte, models.FacebookToken) {
	var token models.FacebookToken

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&token)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
	}

	res, err := http.Get("https://graph.facebook.com/v3.2/me?fields=address,birthday,email,gender,id,location,name&access_token=" + token.Credencial)
	if res.StatusCode != 200 {
		w.WriteHeader(utils.HTTPStatusCode["UNAUTHORIZED"])
		w.Write([]byte("Token inválido"))
		return nil, models.FacebookToken{}
	}
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
		return nil, models.FacebookToken{}
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
		return nil, models.FacebookToken{}
	}

	return body, token
}
