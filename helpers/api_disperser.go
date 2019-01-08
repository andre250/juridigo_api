package helpers

import (
	"bytes"
	"net/http"

	"github.com/juridigo/juridigo_api_usuario/models"
)

/*
Função responsável por controle de middleware
*/
func api(patch string, handleFunction http.HandlerFunc, auth bool) {
	if auth {
		http.HandleFunc(patch, CheckToken(handleFunction))
	} else {
		http.HandleFunc(patch, Cors(handleFunction))
	}
}

/*
APIDisperser função responsável por gerar rotas para uma mesmo mainPath
*/
func APIDisperser(mainPath string, handlerController ...models.DefaultAPI) {
	for _, hc := range handlerController {
		var buffer bytes.Buffer
		buffer.WriteString(mainPath)
		buffer.WriteString(hc.SubPath)
		api(buffer.String(), hc.Handler, hc.Auth)
	}
}
