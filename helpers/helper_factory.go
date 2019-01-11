package helpers

import (
	"github.com/juridigo/juridigo_api_usuario/config"
	"github.com/juridigo/juridigo_api_usuario/models"
)

var configuration models.Config

/*
InitiConfig - Inicializador de configurações
*/
func InitConfig() {

	configuration = config.GetConfig()
}
