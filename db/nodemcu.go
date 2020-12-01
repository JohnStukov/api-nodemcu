package db

import (
	"context"
	"github.com/Ignis-Divine/api-nodemcu/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

	var dbN = MongoCN.Database("nosql")
	var colN = dbN.Collection("nodemcu")

//RevisarSiExisteNodemcu recibe el email de parametro y revisa si ya existe en la db
func RevisarSiExisteNodemcu(tok string) (models.Nodemcu, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	condicion := bson.M{"token":tok,"status":"ok"}
	var resultado models.Nodemcu
	err := colN.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.T
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
