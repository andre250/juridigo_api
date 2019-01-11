package helpers

import (
	"net/http"

	"github.com/rs/cors"
)

/*
Cors - utilizado para comunicação entre containers e local
*/
func Cors(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "DELETE"},
			AllowCredentials: true,
		})
		w.Header().Set("Content-Type", "application/json")
		c.ServeHTTP(w, r, next)
	}
}
