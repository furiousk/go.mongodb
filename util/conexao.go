package util

import (
	"time"

	utils "github.com/furiousk/go.utils/utils"
	mgo "gopkg.in/mgo.v2"
)

//Conn ....
func Conn() (session *mgo.Session, db string) {

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
