package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/juridigo/juridigo_api/config"
	"github.com/juridigo/juridigo_api/models"
	"github.com/juridigo/juridigo_api/routes"
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
	// Obtenção das configurações de ambiente
	configGlobal = config.GetConfig()
	// inicialização da conexão
	// helpers.ConDB()
	// Inicialização do servidor
	serverConfig := []string{":", configGlobal.App.Port}
	fmt.Printf("Mica [User] v%s ouvindo porta: %s", configGlobal.Version, configGlobal.App.Port)
	if http.ListenAndServe(strings.Join(serverConfig, ""), nil) != nil {
		fmt.Println("Porta já esta sendo utilizada")
		log.Fatal()
	}
}
