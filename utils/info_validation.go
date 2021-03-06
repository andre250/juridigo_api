package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/juridigo/juridigo_api_usuario/models"
)

/*
ValidateBasicInfo - Responsável por validar informações basica de cadastro
*/
func ValidateBasicInfo(w http.ResponseWriter, user models.Registro) {
	var erros []models.ErroItem
	errStatus := DateValidation(user.Cadastrais.DataNascimento)
	if errStatus != 0 {
		erros = append(erros, models.ErroItem{
			Msg:  "Data de nascimento invalida",
			Erro: "dataNascimento",
		})
	}

	err := EmailValidation(user.Cadastrais.Email)
	if !err {
		erros = append(erros, models.ErroItem{
			Msg:  "Email inválido",
			Erro: "email",
		})
	}
	err = RgValidation(user.Cadastrais.RG)
	if !err {
		erros = append(erros, models.ErroItem{
			Msg:  "Rg inválido",
			Erro: "rg",
		})
	}

	err = CpfValidation(user.Cadastrais.CPF)
	if !err {
		erros = append(erros, models.ErroItem{
			Msg:  "Cpf inválido",
			Erro: "cpf",
		})
	}

	if len(erros) != 0 {
		listError := models.ErroList{Erros: erros}
		j, _ := json.Marshal(listError)
		w.WriteHeader(HTTPStatusCode["BAD_REQUEST"])
		w.Write(j)
		return
	}
}

/*
ValidatePaymentInfo - Responsǘel por valida informações básicas de pagamento
*/
func ValidatePaymentInfo(w http.ResponseWriter, payment string) (models.Pagamento, error) {
	paymentModel := models.Pagamento{}
	err := json.Unmarshal([]byte(payment), &paymentModel)
	if err != nil {
		w.WriteHeader(HTTPStatusCode["BAD_REQUEST"])
		w.Write([]byte(`{"msg": "Erro no formato do Pagamento", "erro": "pagamento"}`))
		return models.Pagamento{}, errors.New("Erro no formato do pagamaento")
	}

	return paymentModel, nil
}
