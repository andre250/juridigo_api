package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/juridigo/juridigo_api_usuario/models"
	"github.com/juridigo/juridigo_api_usuario/utils"
)

/*
GetId - Recupera dado de IDuser do Facebook
*/
func GetId(w http.ResponseWriter, token string) (string, error) {
	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + token)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro interno na verificação do facebook"))
		return "", errors.New("Erro interno")
	}
	if resp.StatusCode == 400 {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"error":"Token invalido"}`))
		return "", errors.New("Token Invalido")
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro interno na verificação do facebook"))
		return "", errors.New("Erro interno")
	}
	bodyReq := models.FacebookReturn{}
	json.Unmarshal(body, &bodyReq)

	return bodyReq.ID, nil
}
