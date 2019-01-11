package config

import (
	"os"

	"github.com/juridigo/juridigo_api_usuario/models"
)

/*
Prod - Responsável pode difinir confirgurações de ambiente
*/
func prod() {

	globaConfig = models.Config{
		App: models.App{
			Port:   os.Getenv("APP_PORT"),
			Secret: os.Getenv("APP_SECRET"),
		},
		Version: os.Getenv("VER"),
		Database: models.Database{
			Path:     os.Getenv("DB_PATH"),
			Password: os.Getenv("DB_PASS"),
			User:     os.Getenv("DB_USER"),
			Database: os.Getenv("DB_NAME"),
		},
		Amazon: models.Amazon{
			Prefix:  os.Getenv("AWS_FILE_PREFIX"),
			Profile: os.Getenv("AWS_PROFILE"),
			Bucket:  os.Getenv("AWS_BUCKET"),
			Region:  os.Getenv("AWS_REGION"),
		},
		Facebook: models.Facebook{
			AppToken: os.Getenv("FACE_TOKEN"),
		},
	}
}
