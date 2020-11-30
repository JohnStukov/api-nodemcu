package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Datos struct {
	ID         	primitive.ObjectID 	`bson:"_id" json:"id"`
	H   		float64        	`bson:"humedad" json:"humedad"`
	T			float64        	`bson:"temperatura" json:"temperatura"`
	Fecha		string		`bson:"fecha" json:"fecha"`
	Hora		string		`bson:"hora" json:"hora"`
	MacNodemcu	string		`bson:"macNodemcu" json:"macNodemcu"`
}
