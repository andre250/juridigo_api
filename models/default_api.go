package models

import "net/http"

/*
DefaultAPI - Função responsável por definir padrão para argumentos de API
*/
type DefaultAPI struct {
	SubPath string           `json:"subPath"`
	Handler http.HandlerFunc `json:"handler"`
	Auth    bool             `json:"auth"`
}
