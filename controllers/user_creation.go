package controllers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	creditcard "github.com/durango/go-credit-card"
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
	validateBasicInfo(w, registro)

	configuration := config.GetConfig()
	paymentInfo, err := helpers.Decrypt([]byte(string(configuration.App.Secret)), registro.Pagamento)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["UNAUTHORIZED"])
		w.Write([]byte(`{"erro":"Hash de pagamento inválido"}`))
		return
	}

	payment, err := validatePaymentInfo(w, paymentInfo)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Cartão inválido", "erro": "cartão"}`))
		return
	}

	dec, err := base64.StdEncoding.DecodeString(registro.Curriculares.Curriculum)

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte(`{"msg": "Erro ao processar currilum", "erro": "curriculum"}`))
		return
	}

	f, err := os.Create(registro.Credenciais.Credencial + ".pdf")

	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte(`{"msg": "Erro ao processar currilum", "erro": "curriculum"}`))
		return
	}
	f.Write(dec)
	f.Close()
	helpers.AWS().UploadFileToS3(f.Name())
	os.Remove(f.Name())

	registro.Curriculares.Curriculum = "https://s3.amazonaws.com/" + configuration.Amazon.Bucket + "/" + f.Name()

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

	err = helpers.Db().Insert("credenciais", credencial)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte(`{"msg": "Inserção no banco falhou", "erro": "Insert"}`))
		return
	}

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Conta criada com sucesso!"}`))
}

func validatePaymentInfo(w http.ResponseWriter, payment string) (models.Pagamento, error) {
	paymentModel := models.Pagamento{}
	err := json.Unmarshal([]byte(payment), &paymentModel)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Erro no formato do Pagamento", "erro": "pagamento"}`))
		return models.Pagamento{}, errors.New("Erro no formato do pagamaento")
	}

	card := creditcard.Card{Number: paymentModel.Numero, Cvv: paymentModel.Cvv, Month: paymentModel.MesVencimento, Year: paymentModel.AnoVencimento}
	company, err := card.MethodValidate()
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Cartão invalido", "erro": "cartao"}`))
		return models.Pagamento{}, errors.New("Cartao inválido")
	}
	paymentModel.Compania = company.Long

	err = card.Validate()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Cartão invalido", "erro": "cartao"}`))
		return models.Pagamento{}, errors.New("Cartão inválido")
	}
	return paymentModel, nil
}

func validateBasicInfo(w http.ResponseWriter, user models.Registro) {
	var erros []models.ErroItem
	errStatus := utils.DateValidation(user.Cadastrais.DataNascimento)
	if errStatus != 0 {
		erros = append(erros, models.ErroItem{
			Msg:  "Data de nascimento invalida",
			Erro: "dataNascimento",
		})
	}
	err := utils.EmailValidation(user.Cadastrais.Email)
	if !err {
		erros = append(erros, models.ErroItem{
			Msg:  "Email inválido",
			Erro: "email",
		})
	}
	err = utils.RgValidation(user.Cadastrais.RG)
	if !err {
		erros = append(erros, models.ErroItem{
			Msg:  "Rg inválido",
			Erro: "rg",
		})
	}

	err = utils.CpfValidation(user.Cadastrais.CPF)
	if !err {
		erros = append(erros, models.ErroItem{
			Msg:  "Cpf inválido",
			Erro: "cpf",
		})
	}

	if len(erros) != 0 {
		listError := models.ErroList{Erros: erros}
		j, _ := json.Marshal(listError)
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write(j)
		return
	}
}
