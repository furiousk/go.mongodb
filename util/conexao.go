package util

import (
	"strings"
	"time"

	utils "github.com/furiousk/go.utils/utils"
	mgo "gopkg.in/mgo.v2"
)

//Conn faz a conexão em procução
func Conn() (session *mgo.Session, db string) {

	config := utils.ReadOptions(".env")

	_host := config["DB_HOST"]
	_port := config["DB_PORT"]
	_mgdb := config["DB_DATABASE"]
	_user := config["DB_USERNAME"]
	_pwsd := config["DB_PASSWORD"]

	db = config["DB_DATABASE"]

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
func ConnLocal() (session *mgo.Session, db string) {

	config := utils.ReadOptions(".env")

	_host := config["DB_HOST"] + ":" + config["DB_PORT"]
	_mgdb := config["DB_DATABASE"]

	db = config["DB_DATABASE"]

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

	session, db := ConnLocal()
	session.SetMode(mgo.Monotonic, true)
	c = session.DB(db).C(collection)

	return
}
