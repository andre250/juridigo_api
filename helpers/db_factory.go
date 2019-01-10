package helpers

import (
	"errors"
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
	fmt.Println("# Conexão no banco de dados feita com sucesso!")
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

/*
Find - Função de Select CRUD
*/
func (s *Session) Find(collection string, query interface{}, tipo int) (interface{}, error) {
	var result interface{}
	var err error
	queryFunc := s.Session.DB(configuration.Database.Database).C(collection).Find(query)

	if tipo < 0 {
		err = queryFunc.All(&result)
	} else {
		err = queryFunc.One(&result)
	}

	if err != nil {
		return nil, errors.New("Erro ao consultar banco")
	}
	return result, nil
}

/*
FindSelect - Função de insert CRUD
*/
func (s *Session) FindSelect(collection string, query, selector interface{}) {
	var result interface{}
	err := s.Session.DB(configuration.Database.Database).C(collection).Find(query).Select(selector).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

/*
Remove - Função de delete CRUD
*/
func (s *Session) Remove(collection string, query interface{}) {
	err := s.Session.DB(configuration.Database.Database).C(collection).Remove(query)
	if err != nil {
		return
	}
	return

}

/*
Update - Função de update CRUD
*/
func (s *Session) Update(collection string, reference, query interface{}) {
	err := s.Session.DB(configuration.Database.Database).C(collection).Update(reference, query)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
