package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Datos struct {
	ID         	primitive.ObjectID 	`bson:"_id" json:"id"`
	B1   		float64     `bson:"B1" json:"B1"`
	B2   		float64     `bson:"B2" json:"B2"`
	B3   		float64     `bson:"B3" json:"B3"`
	B4   		float64     `bson:"B4" json:"B4"`
	Fecha		string		`bson:"fecha" json:"fecha"`
	Hora		string		`bson:"hora" json:"hora"`
	MacNodemcu	string		`bson:"macNodemcu" json:"macNodemcu"`
}
