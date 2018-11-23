package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Pessoa struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Image     string        `bson:"image"`
	Nome      string        `bson:"nome"`
	Email     string        `bson:"email"`
	Matricula float64       `bson:"matricula"`
	Status    bool          `bson:"status"`
	CreatedAt time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"updatedAt,omitempty"`
}
