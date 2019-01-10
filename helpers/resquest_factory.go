package helpers

import (
	"errors"
	"net/http"

	"github.com/juridigo/juridigo_api_usuario/utils"
)

/*
ReqRefuse - Função de recusa de metodos
*/
func ReqRefuse(w http.ResponseWriter, r *http.Request, methods ...interface{}) error {
	for _, method := range methods {
		if method == r.Method {
			return nil
		}
	}
	w.WriteHeader(utils.HTTPStatusCode["METHOD_NOT_ALLOWED"])
	w.Write([]byte("Metodo não existe"))
	return errors.New("Metodo não existe")
}
