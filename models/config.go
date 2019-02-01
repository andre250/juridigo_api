package models

/*
App - Modelo para inicialização do App
*/
type App struct {
	Port   string `json:"port"`
	Secret string `json:"secret"`
}

/*
Database - Modelo para inicializaçao do Database
*/
type Database struct {
	Path     string `json:"path"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

/*
Email  - Modelo para inicializacao do email
*/
type Email struct {
	From     string `json:"from"`
	Pass     string `json:"pass"`
	SMTP     string `json:"smtp"`
	SMTPAddr string `json:"smtpAddr"`
}

/*
Amazon - Modelo para inicialização da Amazon
*/
type Amazon struct {
	Profile string `json:"profile"`
	Prefix  string `json:"prefix"`
	Bucket  string `json:"bucket"`
	Region  string `json:"region"`
}

/*
Facebook - Modelo para inicialização do Facebook
*/
type Facebook struct {
	AppToken string `json:"appToken"`
}

/*
Config - Model responsável por controlar configurações do Microsserviço
*/
type Config struct {
	Version  string   `json:"version"`
	App      App      `json:"app"`
	Database Database `json:"database"`
	Amazon   Amazon   `json:"amazon"`
	Facebook Facebook `json:"appToken"`
	Email    Email    `json:"email"`
}
