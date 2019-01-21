package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//AlocacaoIntegracao ....
type AlocacaoIntegracao struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Guia      string        `bson:"numero_guia"`
	Matricula string        `bson:"mac_motorista"`
	Nome      string        `bson:"nome_motorista"`
	Pegada    string        `bson:"pegada_motorista"`
	Largada   string        `bson:"largada_motorista"`
	Carro     string        `bson:"numero_carro"`
	Linha     string        `bson:"numero_linha"`
	Situacao  string        `bson:"situacao"`
	Data      time.Time     `bson:"evento,omitempty"`
	CreatedAt time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt time.Time     `bson:"updatedAt,omitempty"`
}
