package helpers

import (
	"fmt"
	"regexp"

	"github.com/juridigo/juridigo_api_usuario/config"
	mgo "gopkg.in/mgo.v2"
)

var mainSession *mgo.Session

/*
Session - Modelo de sessão
*/
type Session struct {
	Session *mgo.Session
}

/*
Connection - Responsável por abir conexão com o bancode Dados
*/
func Connection() {
	configuration = config.GetConfig()
	path := configuration.Database.Path
	path = regexp.MustCompile(`(?m)\<dbuser\>`).ReplaceAllString(path, configuration.Database.User)
	path = regexp.MustCompile(`(?m)\<dbpassword\>`).ReplaceAllString(path, configuration.Database.Password)
	session, err := mgo.Dial(path)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	fmt.Println("# Conexão ao banco de dados feita com sucesso!")
	mainSession = session
}

/*
Db - Função de chamada do bancod
*/
func Db() *Session {
	session := Session{
		Session: mainSession,
	}
	return &session
}

/*
Insert - Função de insert CRUD
*/
func (s *Session) Insert(collection string, inserts interface{}) error {
	err := s.Session.DB(configuration.Database.Database).C(collection).Insert(&inserts)
	if err != nil {
		return err
	}
	return nil
}
