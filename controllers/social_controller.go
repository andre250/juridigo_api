package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juridigo/juridigo_api_usuario/helpers"
	"github.com/juridigo/juridigo_api_usuario/utils"
	"gopkg.in/mgo.v2/bson"

	"github.com/juridigo/juridigo_api_usuario/models"
)

/*
FacebookAuth - Função de autorização social
*/
func FacebookAuth(w http.ResponseWriter, r *http.Request) {

	body, token := facebookInfo(w, r)
	if body == nil {
		return
	}
	bodyReq := models.FacebookReturn{}
	json.Unmarshal(body, &bodyReq)

	query, err := helpers.Db().Find("credenciais", bson.M{"facebookId": bodyReq.ID}, 1)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados de login"))
		return
	}
	json, _ := bson.MarshalJSON(query)
	var credencial models.Credencial
	bson.UnmarshalJSON(json, &credencial)

	if bodyReq.ID == credencial.FacebookID {
		updateLogin(w, credencial, bodyReq, token)
	} else {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte("Dados de login irregulares"))
		return
	}
}

func updateLogin(w http.ResponseWriter, credencial models.Credencial, bodyReq models.FacebookReturn, token models.FacebookToken) {
	if credencial.Credencial != token.Credencial {
		helpers.Db().Update("credenciais", bson.M{"facebookId": bodyReq.ID}, bson.M{"$set": bson.M{"credencial": token.Credencial}})
		fmt.Println("oi")
	}

}
