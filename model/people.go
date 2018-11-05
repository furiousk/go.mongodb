package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Pessoa struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Nome      string        `bson:"name"`
	Email     string        `bson:"email"`
	Matricula float64       `bson:"matricula"`
	Status    bool          `bson:"status"`
	CreatedAt time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"updatedAt,omitempty"`
}
