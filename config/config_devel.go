package config

import (
	"github.com/juridigo/juridigo_api_usuario/models"
)

/*
Devel - Responsável pode difinir confirgurações de ambiente de desenvolvimento
*/
func devel() {
	globaConfig = models.Config{
		App: models.App{
			Port:   "3032",
			Secret: "jur1d1g0",
		},
		Version: "0.0.1",
		Database: models.Database{
			Path:     "mongodb://<dbuser>:<dbpassword>@ds257314.mlab.com:57314/juridevel",
			User:     "juridigo",
			Password: "jur1digo",
			Database: "juridevel",
		},
		Amazon: models.Amazon{
			Prefix:  "juri_",
			Profile: "default",
			Bucket:  "juridigo",
			Region:  "us-east-1",
		},
		Facebook: models.Facebook{
			AppToken: "199329427332694|PluJMN9CqiFsZCVmgNIOA_Z6H08",
		},
		Email: models.Email{
			From:     "gui.martinscaruso@gmail.com",
			Pass:     "Htv1403M@go01338",
			SMTP:     "smtp.gmail.com",
			SMTPAddr: "587",
		},
	}
}
