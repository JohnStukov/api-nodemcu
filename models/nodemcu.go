package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Nodemcu struct {
	ID		primitive.ObjectID `bson:"_id" json:"id"`
	D   	string        `bson:"dispositivo" json:"dispositivo"`
	T		string        `bson:"token" json:"token"`
	S		string        `bson:"status" json:"status"`
}