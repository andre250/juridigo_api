package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/juridigo/juridigo_api_usuario/helpers"
	"github.com/juridigo/juridigo_api_usuario/models"
	"github.com/juridigo/juridigo_api_usuario/utils"
	"gopkg.in/mgo.v2/bson"
)

/*
Login - Função de acesso
*/
func LoginAuth(w http.ResponseWriter, r *http.Request) {
	if helpers.ReqRefuse(w, r, "POST") != nil {
		return
	}
	var credencial models.Login
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&credencial)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
	}

	query, err := helpers.Db().Find("credenciais", bson.M{"credencial": credencial.Credencial}, 1)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte("Usuario inexistente"))
		return
	}

	json, _ := bson.MarshalJSON(query)
	var resultado models.Credencial
	bson.UnmarshalJSON(json, &resultado)

	if resultado.Tipo == 0 {
		defaultAuth(w, resultado)
	} else {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
		return
	}

}

func defaultAuth(w http.ResponseWriter, resultado models.Credencial) {
	query, err := helpers.Db().Find("usuarios", bson.M{"_id": bson.ObjectIdHex(resultado.ID)}, 1)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados de login"))
		return
	}

	json, _ := bson.MarshalJSON(query)
	var user models.Usuario
	bson.UnmarshalJSON(json, &user)

	token := helpers.GenerateLoginToken(resultado.ID, user.Cadastrais.Nome)

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"token":"` + token + `"}`))
	return
}
