package db

import (
	"context"
	"fmt"
	"github.com/Ignis-Divine/api-nodemcu/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var dbA = MongoCN.Database("nosql")
var colA = dbA.Collection("alarmas")

/*InsertoDatos es la parada final con la BD para insertar los datos del usuario */
func InsertoAlarma(alarma models.Alarma) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	alarma.ID = primitive.NewObjectID()
	t := time.Now()
	fec :=""+fmt.Sprintf("%d-%02d-%02d",t.Year(), t.Month(), t.Day())
	hor := ""+fmt.Sprintf("%02d:%02d:%02d",t.Hour(), t.Minute(), t.Second())
	alarma.Fecha=fec
	alarma.Hora=hor
	result, err := colA.InsertOne(ctx, alarma)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

/*
ObtenerRegistros lista todos los usuarios registrados en el sistema,
*/
func ObtenerAlarmas(ID string, page int64, fecha string, hora string, tipo string) ([]*models.Alarma, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var resultados []*models.Alarma
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)
	qury := bson.M{
		//?i no se fija si son mayusculas y minusculas
		"fecha": bson.M{"$regex": `(?i)` + fecha},
		"hora": bson.M{"$regex": `(?i)` + hora},
		"tipo": bson.M{"$regex": `(?i)` + tipo},
	}
	cursor, err := colA.Find(ctx, qury, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	for cursor.Next(ctx) {
		var s models.Alarma
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return resultados, false
		}
		resultados = append(resultados, &s)
	}
	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	cursor.Close(ctx)
	return resultados, false
}
