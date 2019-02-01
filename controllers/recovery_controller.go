package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/juridigo/juridigo_api_usuario/helpers"
	"github.com/juridigo/juridigo_api_usuario/models"
	"github.com/juridigo/juridigo_api_usuario/utils"
)

// RecoveryPassword - Utilizado para gerar token de confirmação de redefinição de senha
func RecoveryPassword(w http.ResponseWriter, r *http.Request) {
	var emailInfo models.Recovery

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&emailInfo)

	_, err := helpers.Db().Find("usuarios", bson.M{
		"cadastrais.email": emailInfo.Email,
		"recoveryToken":    emailInfo.Token,
	}, 1)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Token inválido para esse email"}`))
		return
	}

	newToken := utils.GenerateRecoveryToken(4)

	helpers.Db().Update("usuarios", bson.M{"cadastrais.email": emailInfo.Email}, bson.M{"$set": bson.M{"recoveryToken": newToken}})

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"token": "` + newToken + `"}`))
}

//ChangePassword - Redefinição de senha
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var newInfo models.ChangePass

	if newInfo.Recovery.Token == "" {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Token inválido para esse email"}`))
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newInfo)

	usuariofinal := models.Usuario{}
	err := helpers.Db().FindSelectUser("usuarios", bson.M{
		"cadastrais.email": newInfo.Recovery.Email,
		"recoveryToken":    newInfo.Recovery.Token,
	}, &usuariofinal)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Token inválido para esse email"}`))
		return
	}
	userID := usuariofinal.ID.Hex()

	usuarioCredenciais, err := helpers.Db().Find("credenciais", bson.M{
		"id": userID,
	}, 1)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Erro de parse"}`))
		return
	}

	credencial := models.Credencial{}

	bytesCred, _ := json.Marshal(usuarioCredenciais)
	json.Unmarshal(bytesCred, &credencial)

	bytes := sha256.Sum256([]byte(credencial.UsuarioDetalhe + "@" + newInfo.NewPassword))
	sha := hex.EncodeToString(bytes[:])

	if credencial.Tipo == 0 {
		helpers.Db().Update("credenciais", bson.M{"id": userID}, bson.M{"$set": bson.M{"credencial": sha}})
	} else if credencial.Tipo == 1 {
		helpers.Db().Update("credenciais", bson.M{"id": userID}, bson.M{"$set": bson.M{"recuperacaoLogin": sha}})
	}

	helpers.Db().Update("usuarios", bson.M{"_id": bson.ObjectIdHex(userID)}, bson.M{"$set": bson.M{"recoveryToken": ""}})

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Senha atualizada com sucesso!"}`))
}
