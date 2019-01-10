package models

/*
Login - Padrão de chamada de login
*/
type Login struct {
	Credencial string `bson:"credencial" json:"credencial"`
}
