package routes

import (
	"fmt"
	"net/http"

	"github.com/juridigo/juridigo_api/controllers"
	"github.com/juridigo/juridigo_api/helpers"
	"github.com/juridigo/juridigo_api/models"
)

/*
Routes - Controlador de rotas do microsserviço
*/
func Routes() {
	helpers.APIDisperser("/user",
		models.DefaultAPI{SubPath: "/register", Handler: controllers.CreateUser, Auth: false},
	)
}

func nova(w http.ResponseWriter, r *http.Request) {
	fmt.Println("oi")
}
