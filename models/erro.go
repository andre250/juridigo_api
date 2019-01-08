package models

type ErroList struct {
	Erros []ErroItem `json:"erros"`
}

type ErroItem struct {
	Msg  string `json:"msg"`
	Erro string `json:"erro"`
}
