package config

import "github.com/juridigo/juridigo_api/models"

/*
Devel - Responsável pode difinir confirgurações de ambiente de desenvolvimento
*/
func devel() {
	globaConfig = models.Config{
		App: models.App{
			Port:   "3030",
			Secret: "JUR1d1G00S3cr377",
		},
		Version: "0.0.1",
		Database: models.Database{
			Path:     "mongodb://<dbuser>:<dbpassword>@ds229701.mlab.com:29701/mica-develop",
			User:     "micauser",
			Password: "micapass1",
		},
	}

}
