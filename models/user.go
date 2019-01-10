package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*
Usuario - Modelo de inicialização de um usuario
*/
type Usuario struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	Cadastrais     Cadastrais    `bson:"cadastrais" json:"cadastrais"`
	Curriculares   Curriculares  `bson:"curriculares" json:"curriculares"`
	DadosPagamento Pagamento     `bson:"DadosPagamento" json:"DadosPagamento"`
}

/*
Registro - Controlador de registro
*/
type Registro struct {
	Credenciais  Credencial   `bson:"credenciais" json:"credenciais"`
	Cadastrais   Cadastrais   `bson:"cadastrais" json:"cadastrais"`
	Curriculares Curriculares `bson:"curriculares" json:"curriculares"`
	Pagamento    string       `bson:"pagamento" json:"pagamento"`
}

/*
Credencial - Controlador de acesso
*/
type Credencial struct {
	ID         string `bson:"id" json:"id"`
	Credencial string `bson:"credencial" json:"credencial"`
	Tipo       int    `bson:"tipo" json:"tipo"`
	FacebookID string `bson:"facebookId,omitempty" json:"facebookId"`
}

/*
Cadastrais - Modelo de inicialização de dados cadastrais de um model.Usuario
*/
type Cadastrais struct {
	Nome           string    `bson:"nome" json:"nome"`
	DataNascimento time.Time `bson:"dataNascimento" json:"dataNascimento"`
	Email          string    `bson:"email" json:"email"`
	Telefone       string    `bson:"telefone" json:"telefone"`
	RG             string    `bson:"rg" json:"rg"`
	CPF            string    `bson:"cpf" json:"cpf"`
	CEP            string    `bson:"cep" json:"cep"`
	Cidade         string    `bson:"cidade" json:"cidade"`
	Bairro         string    `bson:"bairro" json:"bairro"`
	Rua            string    `bson:"rua" json:"rua"`
	Número         string    `bson:"numero" json:"numero"`
	Complemento    string    `bson:"complemento" json:"complemento"`
	Longitude      float64   `bson:"longitude" json:"longitude"`
	Latitude       float64   `bson:"latitude" json:"latitude"`
}

/*
Curriculares - Modelo de inicialização de dados curriculares de um model.Usuario
*/
type Curriculares struct {
	Formacao   []Formacao `bson:"formacao" json:"formacao"`
	OAB        string     `bson:"oab" json:"oab"`
	Curriculum string     `bson:"curriculum" bson:"curriculum"`
}

/*
Formacao - Modelo de inicializaçã de formações de um model.Curriculares
*/
type Formacao struct {
	Escolaridade int    `bson:"escolaridade" json:"escolaridade"`
	Instituicao  string `bson:"instituicao" json:"instituicao"`
	AnoConclusao string `bson:"anoConclusao" json:"anoConclusao"`
}

/*
Pagamento - Modelo de inicialização de dados de pagamento de um model.Usuario
*/
type Pagamento struct {
	Compania      string `bson:"compania" json:"companhia"`
	Numero        string `bson:"numero" json:"numero"`
	Banco         string `bson:"banco" json:"banco"`
	Cvv           string `bson:"cvv" json:"cvv"`
	AnoVencimento string `bson:"anoVencimento" json:"anoVencimento"`
	MesVencimento string `bson:"mesVencimento" json:"mesVencimento"`
	Agencia       string `bson:"agencia" json:"agencia"`
	Conta         string `bson:"conta" json:"conta"`
}
