package controllers

import (
	"encoding/json"
	"net/http"
	"net/smtp"

	"github.com/juridigo/juridigo_api_usuario/helpers"
	"gopkg.in/mgo.v2/bson"

	"github.com/juridigo/juridigo_api_usuario/utils"

	"github.com/juridigo/juridigo_api_usuario/models"

	"github.com/juridigo/juridigo_api_usuario/config"
)

/*
SendEmail - Responsável por disparar emails em geral
*/
func SendEmail(w http.ResponseWriter, r *http.Request) {
	if helpers.ReqRefuse(w, r, "POST") != nil {
		return
	}

	var emailInfo models.SendEmail

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&emailInfo)

	configuration := config.GetConfig()

	from := configuration.Email.From
	pass := configuration.Email.Pass
	to := emailInfo.To

	switch emailInfo.Type {
	case 1:
		recoveryEmail(w, from, pass, to, emailInfo.Name, configuration.Email.SMTP, configuration.Email.SMTPAddr)
	}

}

func recoveryEmail(w http.ResponseWriter, from, pass, to, name, smtpURL, smtpAddr string) {

	_, err := helpers.Db().Find("usuarios", bson.M{"cadastrais.email": to}, 1)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["NOT_FOUND"])
		w.Write([]byte(`{"msg": "Email não cadastrado."}`))
		return
	}

	token := utils.GenerateRecoveryToken(5)

	helpers.Db().Update("usuarios", bson.M{"cadastrais.email": to}, bson.M{"$set": bson.M{"recoveryToken": token}})

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: [JURIDIGO] Recuperacao de senha \n\n" +
		"Ola " + name + ". Seu token de confirmaçao: " + token

	err = smtp.SendMail(smtpURL+":"+smtpAddr,
		smtp.PlainAuth("", from, pass, smtpURL),
		from, []string{to}, []byte(msg))

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte(`{"msg": "Algo errado do nosso lado."}`))
	}

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Email enviado com sucesso!"}`))
}
