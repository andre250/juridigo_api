package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juridigo/juridigo_api_usuario/models"
	"github.com/juridigo/juridigo_api_usuario/utils"
)

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
