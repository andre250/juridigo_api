package models

/*
FacebookToken - Busca de token
*/
type FacebookToken struct {
	Credencial string `bson:"credencial" json:"credencial"`
}

/*
FacebookReturn - Retorno de busca do token
*/
type FacebookReturn struct {
	ID              string      `bson:"id" json:"id"`
	Nome            string      `bson:"name" json:"name"`
	Email           string      `bson:"email" json:"email"`
	DataAniversario string      `bson:"birthday" json:"birthday"`
	Genero          string      `bson:"gender" json:"gender"`
	Localização     Localizacao `bson:"location" json:"location"`
}

/*
Localizacao - complement models.FacebookReturn
*/
type Localizacao struct {
	ID   string `bson:"id" json:"id"`
	Nome string `bson:"name" json:"name"`
}
