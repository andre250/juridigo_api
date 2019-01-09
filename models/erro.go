package models

/*
ErroList - Modelo de Apresentação de Erros
*/
type ErroList struct {
	Erros []ErroItem `json:"erros"`
}

/*
ErroItem - Modelo de incremendo do model.ErroList
*/
type ErroItem struct {
	Msg  string `json:"msg"`
	Erro string `json:"erro"`
}
