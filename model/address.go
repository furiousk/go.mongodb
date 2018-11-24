package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Type       string        `bson:"type"`
	Properties struct {
		FormattedAddress string `bson:"formatted_address" required:"true" unique:"true"`
	}
	Status    bool      `bson:"status"`
	CreatedAt time.Time `bson:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
}
