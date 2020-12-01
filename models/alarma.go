package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Alarma struct {
	ID		primitive.ObjectID `bson:"_id" json:"id"`
	Alarma   	string        `bson:"alarma" json:"alarma"`
	Tipo		string        `bson:"tipo" json:"tipo"`
	Fecha		string		`bson:"fecha" json:"fecha"`
	Hora		string		`bson:"hora" json:"hora"`
	MacNodemcu	string		`bson:"macNodemcu" json:"macNodemcu"`
}
