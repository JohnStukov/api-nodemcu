package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Datos struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Movimiento  string             `bson:"movimiento" json:"movimiento"`
	Sonico      float32            `bson:"sonico" json:"sonico"`
	Temperatura float32            `bson:"temperatura" json:"temperatura"`
	Humedad     float32            `bson:"humedad" json:"humedad"`
	Fecha       string             `bson:"fecha" json:"fecha"`
	Hora        string             `bson:"hora" json:"hora"`
	MacNodemcu  string             `bson:"macNodemcu" json:"macNodemcu"`
}
