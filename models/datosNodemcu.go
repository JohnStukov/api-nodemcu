package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Datos struct {
	ID         	primitive.ObjectID 	`bson:"_id" json:"id"`
	B1   		string     `bson:"B1" json:"B1"`
	Est1		string		`bson:"est1" json:"est1"`
	B2   		string     `bson:"B2" json:"B2"`
	Est2		string		`bson:"est2" json:"est2"`
	B3   		string     `bson:"B3" json:"B3"`
	Est3		string		`bson:"est3" json:"est3"`
	B4   		string     `bson:"B4" json:"B4"`
	Est4		string		`bson:"est4" json:"est4"`
	Fecha		string		`bson:"fecha" json:"fecha"`
	Hora		string		`bson:"hora" json:"hora"`
	MacNodemcu	string		`bson:"macNodemcu" json:"macNodemcu"`
}
