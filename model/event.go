package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Status    int           `bson:"status"`
	Nome      string        `bson:"nome"`
	Codigo    int           `bson:"codigo"`
	CreatedAt time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"updatedAt,omitempty"`
}
