package controllers

import (
	"net/http"
	"os"

	"github.com/juridigo/juridigo_api_usuario/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(utils.HTTPStatusCode["OK"])
	w.Write([]byte(`{"msg": "Estou vivo em ` + os.Getenv("ENV") + `}`))
}
