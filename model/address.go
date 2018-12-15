package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	//Props ...
	Props struct {
		FormattedAddress string `bson:"formatted_address" json:"formatted_address"`
	}
	//Address ...
	Address struct {
		ID         bson.ObjectId `bson:"_id,omitempty" json:"_id"`
		Properties Props         `bson:"properties" json:"properties"`
		Type       string        `bson:"type" json:"type"`
		Status     bool          `bson:"status" json:"status"`
		CreatedAt  time.Time     `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
		UpdatedAt  time.Time     `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	}
)
