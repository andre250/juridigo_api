package helpers

import (
	"os"
	"regexp"

	mgo "gopkg.in/mgo.v2"
)

var mainSession *mgo.Session

/*
Connection - Responsável por abir conexão com o bancode Dados
*/
func Connection() {
	path := os.Getenv("DB_PATH")
	regexp.MustCompile("<dbuser>").ReplaceAllString(path, os.Getenv("DB_USER"))
	regexp.MustCompile("<dbpassword>").ReplaceAllString(path, os.Getenv("DB_PASS"))
	session, err := mgo.Dial(path)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	mainSession = session
}

/*
Db - Responsável por obter conexão para execução
*/
func Db() *mgo.Session {
	return mainSession
}
