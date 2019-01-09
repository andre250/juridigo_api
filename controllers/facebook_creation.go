package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/juridigo/juridigo_api_usuario/models"
	"github.com/juridigo/juridigo_api_usuario/utils"
)

/*
GetFacebookInfo - Responsável por retornar ifnormações do front
*/
func GetFacebookInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(utils.HTTPStatusCode["METHOD_NOT_ALLOWED"])
		w.Write([]byte("Metodo não existe"))
	}
	var token models.FacebookToken

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&token)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
	}

	res, err := http.Get("https://graph.facebook.com/v3.2/me?fields=address,birthday,email,gender,id,location,name&access_token=" + token.Token)
	if res.StatusCode != 200 {
		w.WriteHeader(utils.HTTPStatusCode["UNAUTHORIZED"])
		w.Write([]byte("Token inválido"))
		return
	}
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
		return
	}
	bodyReq := models.FacebookReturn{}
	json.Unmarshal(body, &bodyReq)

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write(body)
}

func CreateFacebookUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(utils.HTTPStatusCode["METHOD_NOT_ALLOWED"])
		w.Write([]byte("Metodo não existe"))
	}
	var registro models.Registro

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&registro)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
	}

	fmt.Println(registro.Credenciais.Credencial)
}
