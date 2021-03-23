package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Datos struct {
	ID         	primitive.ObjectID 	`bson:"_id" json:"id"`
	B   		string     `bson:"B" json:"B"`
	Est		string		`bson:"est" json:"est"`
	Fecha		string		`bson:"fecha" json:"fecha"`
	Hora		string		`bson:"hora" json:"hora"`
	MacNodemcu	string		`bson:"macNodemcu" json:"macNodemcu"`
}
