package helpers

import (
	"fmt"
	"regexp"

	"github.com/juridigo/juridigo_api_usuario/config"

	mgo "gopkg.in/mgo.v2"
)

var mainSession *mgo.Session

/*
Connection - Responsável por abir conexão com o bancode Dados
*/
func Connection() {
	config := config.GetConfig()
	path := config.Database.Path
	path = regexp.MustCompile(`(?m)\<dbuser\>`).ReplaceAllString(path, config.Database.User)
	path = regexp.MustCompile(`(?m)\<dbpassword\>`).ReplaceAllString(path, config.Database.Password)
	session, err := mgo.Dial(path)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	fmt.Println("# Conexão ao banco de dados feita com sucesso!")
	mainSession = session
}

/*
Db - Responsável por obter conexão para execução
*/
func Db() *mgo.Session {
	return mainSession
}
