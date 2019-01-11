package models

/*
Login - chamada de login padrão
*/
type Login struct {
	Credencial string `bson:"credencial" json:"credencial"`
}
