package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/juridigo/juridigo_api_usuario/config"
	"github.com/juridigo/juridigo_api_usuario/helpers"
	"github.com/juridigo/juridigo_api_usuario/models"
	"gopkg.in/mgo.v2/bson"

	"github.com/juridigo/juridigo_api_usuario/utils"
)

/*
CreateUser - Função responsável pelo cadastro de um novo usuario
*/
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if helpers.ReqRefuse(w, r, "POST") != nil {
		return
	}
	var registro models.Registro

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&registro)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
	}
	utils.ValidateBasicInfo(w, registro)
	configuration := config.GetConfig()
	paymentInfo, err := helpers.Decrypt([]byte(string(configuration.App.Secret)), registro.Pagamento)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["UNAUTHORIZED"])
		w.Write([]byte(`{"erro":"Hash de pagamento inválido"}`))
		return
	}

	payment, err := utils.ValidatePaymentInfo(w, paymentInfo)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Cartão inválido", "erro": "cartão"}`))
		return
	}

	url := helpers.UploadFile(w, registro.Credenciais.Credencial, registro.Curriculares.Curriculum)

	if url == "" {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte(`{"msg": "Erro ao processar currilum", "erro": "curriculum"}`))
		return
	}
	registro.Curriculares.Curriculum = url

	user := models.Usuario{
		ID:             bson.NewObjectId(),
		Cadastrais:     registro.Cadastrais,
		Curriculares:   registro.Curriculares,
		DadosPagamento: payment,
	}

	err = helpers.Db().Insert("usuarios", user)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte(`{"msg": "Inserção no banco falhou", "erro": "Insert"}`))
		return
	}

	credencial := models.Credencial{
		ID:         strings.Split(strings.Split(user.ID.String(), "ObjectIdHex(\"")[1], "\")")[0],
		Credencial: registro.Credenciais.Credencial,
		Tipo:       registro.Credenciais.Tipo,
	}

	if credencial.Tipo == 1 {
		idUser, err := helpers.GetId(w, credencial.Credencial)
		if err != nil {
			helpers.Db().Remove("usuarios", bson.M{"cadastrais.email": user.Cadastrais.Email})
			return
		}
		credencial.FacebookID = idUser
		credencial.RecuperacaoLogin = registro.Credenciais.RecuperacaoLogin
	}

	err = helpers.Db().Insert("credenciais", credencial)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte(`{"msg": "Inserção no banco falhou", "erro": "Insert"}`))
		return
	}

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Conta criada com sucesso!"}`))
}
