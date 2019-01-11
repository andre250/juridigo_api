package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/juridigo/juridigo_api_usuario/models"
)

var globaConfig models.Config

/*
SetConfig - Função responsavel por controlar variaveis de ambiente do microsserviço
*/
func SetConfig(wg *sync.WaitGroup) {
	if os.Getenv("ENV") == "Production" || os.Getenv("ENV") == "Staging" {
		prod()
		list, err := configValidator()

		if !err {
			fmt.Println("  ⬐ Precisam ser definidos")
			fmt.Println(list)
			return
		}
		if os.Getenv("ENV") == "Production" {
			fmt.Println("# Usando configuração de Produção")
		} else {
			fmt.Println("# Usando configuração de Staging")
		}
	} else {
		devel()
		list, err := configValidator()

		if !err {
			fmt.Println("  ⬐ Precisam ser definidos")
			fmt.Println(list)
			return
		}
		fmt.Println("# Usando configuração de Desenvolvimento")
	}

	wg.Done()
}

/*
GetConfig - Função responsavel por obter variaveis de ambiente do microsserviço
*/
func GetConfig() models.Config {
	return globaConfig
}
