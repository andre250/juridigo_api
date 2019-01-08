package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	creditcard "github.com/durango/go-credit-card"
	"github.com/juridigo/juridigo_api/config"
	"github.com/juridigo/juridigo_api/helpers"
	"github.com/juridigo/juridigo_api/models"

	"github.com/juridigo/juridigo_api/utils"
)

/*
CreateUser - Função responsável pelo cadastro de um novo usuario
*/
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(utils.HTTPStatusCode["METHOD_NOT_ALLOWED"])
		w.Write([]byte("Metodo não existe"))
	}
	var user models.Usuario

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["INTERNAL_SERVER_ERROR"])
		w.Write([]byte("Erro ao obter dados"))
	}
	validateBasicInfo(w, user)

	configuration := config.GetConfig()
	payment, err := helpers.Decrypt([]byte(string(configuration.App.Secret)), user.Pagamento)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["UNAUTHORIZED"])
		w.Write([]byte(`{"erro":"Hash de pagamento inválido"}`))
		return
	}

	validatePaymentInfo(w, payment)
}

func validatePaymentInfo(w http.ResponseWriter, payment string) {
	paymentModel := models.Pagamento{}
	err := json.Unmarshal([]byte(payment), &paymentModel)
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Erro no formato do Pagamento", "erro": "pagamento"}`))
		return
	}

	card := creditcard.Card{Number: paymentModel.Numero, Cvv: paymentModel.Cvv, Month: paymentModel.MesVencimento, Year: paymentModel.AnoVencimento}
	company, err := card.Method()
	if err != nil {
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Cartão invalido", "erro": "cartao"}`))
		return
	}
	paymentModel.Banco = company.Long
	err = card.Validate()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(utils.HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Cartão invalido", "erro": "cartao"}`))
		return
	}

}

func validateBasicInfo(w http.ResponseWriter, user models.Usuario) {
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
