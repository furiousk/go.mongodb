package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Telemetria struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Serial    int           `bson:"serial"`
		Status    int           `bson:"status"`
		Name      string        `bson:"name"`
		Codmix    int           `bson:"codmix"`
		SysDate   time.Time     `bson:"systimestamp"`
		GpsDate   time.Time     `bson:"gpstimestamp"`
		Avl       int           `bson:"backofficelk"`
		Condutor  int           `bson:"condutor"`
		Evento    bson.ObjectId `bson:"evento,omitempty"`
		Position  bson.ObjectId `bson:"position,omitempty"`
		CreatedAt time.Time     `bson:"createdAt,omitempty"`
		UpdatedAt time.Time     `bson:"updatedAt,omitempty"`
	}
)
