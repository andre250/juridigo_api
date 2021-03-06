package routes

import (
	"github.com/juridigo/juridigo_api_usuario/controllers"
	"github.com/juridigo/juridigo_api_usuario/helpers"
	"github.com/juridigo/juridigo_api_usuario/models"
)

/*
Routes - Controlador de rotas do microsserviço
*/
func Routes() {
	helpers.APIDisperser("",
		models.DefaultAPI{SubPath: "/register", Handler: controllers.CreateUser, Auth: false},
		models.DefaultAPI{SubPath: "/facebook", Handler: controllers.GetFacebookInfo, Auth: false},
		models.DefaultAPI{SubPath: "/recovery", Handler: controllers.RecoveryPassword, Auth: false},
		models.DefaultAPI{SubPath: "/newPassword", Handler: controllers.ChangePassword, Auth: false},
		models.DefaultAPI{SubPath: "/login", Handler: controllers.LoginAuth, Auth: false},
		models.DefaultAPI{SubPath: "/login/facebook", Handler: controllers.FacebookAuth, Auth: false},
		models.DefaultAPI{SubPath: "/login/facebook/recovery", Handler: controllers.RecoveryLogin, Auth: false},
		models.DefaultAPI{SubPath: "/email/send", Handler: controllers.SendEmail, Auth: false},
	)

	helpers.APIDisperser("/health",
		models.DefaultAPI{SubPath: "/amilive", Handler: controllers.HealthCheck, Auth: false},
	)
}
