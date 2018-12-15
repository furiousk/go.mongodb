package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	//Geo ....
	Geo struct {
		Type        string     `bson:"type"`
		Coordinates [2]float64 `bson:"coordinates"`
	}
	//Point ....
	Point struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Geometry  Geo           `bson:"geometry"`
		Address   bson.ObjectId `bson:"address"`
		Status    bool          `bson:"status"`
		CreatedAt time.Time     `bson:"createdAt,omitempty"`
		UpdatedAt time.Time     `bson:"updatedAt,omitempty"`
	}
)
