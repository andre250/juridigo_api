package models

// Recovery - Modelo de bypass de recuperação de senha
type Recovery struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// ChangePass - Modelo de resstruturcação de senha
type ChangePass struct {
	NewPassword string   `json:"newPassword"`
	Recovery    Recovery `json:"authInfo"`
}
