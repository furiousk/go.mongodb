package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Evento struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Evento      int           `bson:"evento"`
	Serial      int           `bson:"serial"`
	Ocorrencia  int           `bson:"ocorrencia"`
	Tempototal  int           `bson:"tempototal"`
	Tempoinicio time.Time     `bson:"tempoinicio"`
	Tempofinal  time.Time     `bson:"tempofinal"`
	Condutor    int           `bson:"condutor"`
	Gpsinicio   bson.ObjectId `bson:"gpsinicio,omitempty"`
	Gpsfinal    bson.ObjectId `bson:"gpsfinal,omitempty"`
	Telemetria  bson.ObjectId `bson:"telemetria,omitempty"`
	CreatedAt   time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time     `bson:"updatedAt,omitempty"`
}
