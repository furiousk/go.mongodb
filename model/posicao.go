package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Posicao ....
type Posicao struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	Identifier    string        `bson:"identifier"`
	Geometry      Geo           `bson:"geometry"`
	Speed         float64       `bson:"speed"`
	Course        int           `bson:"course"`
	Odometer      float64       `bson:"odometer"`
	Hourimeter    float64       `bson:"hourimeter"`
	Timestamp     time.Time     `bson:"timestamp"`
	Avl           int           `bson:"backoffice"`
	Altitude      float64       `bson:"altitude"`
	BatteryCharge string        `bson:"battery_charge"`
	BatteryStatus string        `bson:"battery_status"`
	VAccuracy     float64       `bson:"vaccuracy"`
	HAccuracy     float64       `bson:"haccuracy"`
	Network       string        `bson:"network"`
	Model         string        `bson:"model"`
	ModelVersion  string        `bson:"model_version"`
	CreatedAt     time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt     time.Time     `bson:"updatedAt,omitempty"`
}
