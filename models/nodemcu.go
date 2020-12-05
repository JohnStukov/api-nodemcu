package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Nodemcu struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Dispositivo string             `bson:"dispositivo" json:"dispositivo"`
	Token       string             `bson:"token" json:"token"`
	Status      string             `bson:"status" json:"status"`
}