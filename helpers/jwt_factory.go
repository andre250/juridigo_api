package helpers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/juridigo/juridigo_api_usuario/config"
	"github.com/juridigo/juridigo_api_usuario/utils"
)

/*
GenerateLoginToken - Metodo de criação de token JWT
*/
func GenerateLoginToken(id, name, status string, latitude, longitude float64) string {
	config := config.GetConfig()
	var env string
	switch os.Getenv("ENV") {
	case "Production":
		env = "prod"
	case "Staging":
		env = "stag"
	default:
		env = "devel"

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          id,
		"name":        name,
		"environment": env,
		"exp":         time.Now().Add(time.Hour * 8).Unix(),
		"latitude":    latitude,
		"longitude":   longitude,
		"status":      status,
	})

	tokenString, _ := token.SignedString([]byte(config.App.Secret))

	return tokenString
}

/*
CheckToken - Função responsável por controlar acesso a rotas não públicas
*/
func CheckToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		config := config.GetConfig()
		if len(r.Header["Authtoken"]) == 0 {
			w.WriteHeader(utils.HTTPStatusCode["UNAUTHORIZED"])
			w.Write([]byte("Token inexistente"))
			return
		}
		tokenToVerify := r.Header.Get("AuthToken")
		token, err := jwt.Parse(tokenToVerify, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.App.Secret), nil
		})

		var env string

		switch os.Getenv("ENV") {
		case "Production":
			env = "prod"
		case "Staging":
			env = "stag"
		default:
			env = "devel"

		}
		if token.Claims.(jwt.MapClaims)["environment"] != env {
			w.WriteHeader(utils.HTTPStatusCode["UNAUTHORIZED"])
			w.Write([]byte("Wrong environment"))
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				w.WriteHeader(utils.HTTPStatusCode["UNAUTHORIZED"])
				w.Write([]byte("Token inválido"))
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				w.WriteHeader(utils.HTTPStatusCode["UNAUTHORIZED"])
				w.Write([]byte("Token expirado"))
				return
			} else {
				fmt.Println("Couldn't handle this token:", err)
				return
			}
		} else {
			fmt.Println("Couldn't handle this token:", err)
			return
		}
	}
}
