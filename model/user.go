package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//User .....
type User struct {
	ID        bson.ObjectId   `bson:"_id,omitempty"`
	LastLogin time.Time       `bson:"last_login"`
	Perfis    []bson.ObjectId `bson:"perfis"`
	Password  string          `bson:"password"`
	Pessoa    bson.ObjectId   `bson:"pessoa"`
	Status    bool            `bson:"status"`
	CreatedAt time.Time       `bson:"createdAt,omitempty"`
	UpdatedAt time.Time       `bson:"updatedAt,omitempty"`
}
