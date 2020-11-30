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

var dbD = MongoCN.Database("nosql")
var colD = dbD.Collection("prueba")

/*InsertoDatos es la parada final con la BD para insertar los datos del usuario */
func InsertoDatos(datos models.Datos) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	datos.ID = primitive.NewObjectID()
	t := time.Now()
	fec :=""+fmt.Sprintf("%d-%02d-%02d",t.Year(), t.Month(), t.Day())
	hor := ""+fmt.Sprintf("%02d:%02d:%02d",t.Hour(), t.Minute(), t.Second())
	datos.Fecha=fec
	datos.Hora=hor
	result, err := colD.InsertOne(ctx, datos)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

/*
ObtenerRegistros lista todos los usuarios registrados en el sistema,
*/
func ObtenerRegistros(ID string, page int64, fecha string, hora string) ([]*models.Datos, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var resultados []*models.Datos
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)
	qury := bson.M{
		//?i no se fija si son mayusculas y minusculas
		"fecha": bson.M{"$regex": `(?i)` + fecha},
		"hora": bson.M{"$regex": `(?i)` + hora},
	}
	cursor, err := colD.Find(ctx, qury, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	for cursor.Next(ctx) {
		var s models.Datos
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
