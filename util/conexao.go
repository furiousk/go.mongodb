package util

import (
	"log"
	"strings"
	"time"

	"github.com/furiousk/go.models/models"
	utils "github.com/furiousk/go.utils/utils"
	mgo "gopkg.in/mgo.v2"
)

func Connect() (session *mgo.Session, db string) {

	ENVIRONMENT := utils.ReadOptions(".env")

	switch ENVIRONMENT.Appenv {

	case "development":
		return ConnLocal(ENVIRONMENT)

	case "production":
		return Conn(ENVIRONMENT)

	default:
		log.Println("Não encontrado!")
		return session, ""
	}
}

//Conn faz a conexão em procução
func Conn(config *models.Environment) (session *mgo.Session, db string) {

	_host := config.Dbhost
	_port := config.Dbport
	_mgdb := config.Dbdatabase
	_user := config.Dbusername
	_pwsd := config.Dbpassword

	db = config.Dbdatabase

	_p1 := strings.Join([]string{_user, _pwsd}, ":")
	_p2 := strings.Join([]string{_host, _port}, ":")
	_p3 := strings.Join([]string{_mgdb, "authSource=admin"}, "?")

	_userpwsd := strings.Join([]string{_p1, _p2}, "@")

	_url := strings.Join([]string{_userpwsd, _p3}, "/")

	session, err := mgo.Dial(_url)

	utils.Check(err)

	session.SetMode(mgo.Monotonic, true)

	return
}

//ConnLocal faz a conexão ao servidor de desenvolvimento
func ConnLocal(config *models.Environment) (session *mgo.Session, db string) {

	_host := config.Dbhost + ":" + config.Dbport
	_mgdb := config.Dbdatabase

	db = config.Dbdatabase

	mongoDBDialInfo := &mgo.DialInfo{

		Addrs:    []string{_host},
		Timeout:  60 * time.Second,
		Database: _mgdb,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)

	utils.Check(err)

	session.SetMode(mgo.Monotonic, true)

	return
}

func SetCollection(collection string) (c *mgo.Collection, session *mgo.Session) {

	session, db := Connect()
	c = session.DB(db).C(collection)

	return
}
