package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/juridigo/juridigo_api_usuario/config"
	"github.com/juridigo/juridigo_api_usuario/helpers"
	"github.com/juridigo/juridigo_api_usuario/models"
	"github.com/juridigo/juridigo_api_usuario/routes"
)

var wg sync.WaitGroup
var configGlobal models.Config

func main() {
	// Inicialização das rotas
	routes.Routes()

	//Processo de definição das configurações
	wg.Add(1)
	config.SetConfig(&wg)
	wg.Wait()
	helpers.InitConfig()
	helpers.MakeSession()
	// Obtenção das configurações de ambiente
	configGlobal = config.GetConfig()
	// inicialização da conexão
	helpers.Connection()
	// Inicialização do servidor
	serverConfig := []string{":", configGlobal.App.Port}
	fmt.Printf("Juridigo [User] v%s ouvindo porta: %s\n", configGlobal.Version, configGlobal.App.Port)
	if http.ListenAndServe(strings.Join(serverConfig, ""), nil) != nil {
		fmt.Println("Porta já esta sendo utilizada")
		log.Fatal()
	}
}
