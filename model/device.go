package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Device struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Serial    int32         `bson:"serial"`
	Version   int32         `bson:"version"`
	Product   int32         `bson:"product"`
	Active    bool          `bson:"active"`
	CreatedAt time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"updatedAt,omitempty"`
}
