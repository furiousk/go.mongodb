package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Point struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Geometry struct {
		Type        string    `bson:"type"`
		Coordinates []float64 `bson:"coordinates"`
	}
	Address   bson.ObjectId `bson:"address"`
	Status    bool          `bson:"status"`
	CreatedAt time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"updatedAt,omitempty"`
}
