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

	var dbN = MongoCN.Database("nosql")
	var colN = dbN.Collection("nodemcu")

/*InsertoRegistroUsuario es la parada final con la BD para insertar los datos del usuario */
func InsertoRegistroNodemcu(n models.Nodemcu) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	n.ID = primitive.NewObjectID()
	result, err := colN.InsertOne(ctx, n)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

func BorroNodemcu(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}
	_, err := colN.DeleteOne(ctx, condicion)
	return err
}

//ModificoRegistro permite modificar los datos del usuario
func ModificoNodemcu(n models.Nodemcu, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	//armar el registro de la actualizacion de la db // Ejemplo 1
	registro := make(map[string]interface{})
	if len(n.Dispositivo) > 0 {
		registro["dispositivo"] = n.Dispositivo
	}
	if len(n.Token) > 0 {
		registro["token"] = n.Token
	}
	if len(n.Status) > 0 {
		registro["status"] = n.Status
	}
	updateString := bson.M{
		"$set": registro,
	}
	//Proceso de actualizacion
	//convierte el ID de string a objID
	objID, _ := primitive.ObjectIDFromHex(ID)
	//filtar nodemcu
	filtro := bson.M{"_id": bson.M{"$eq": objID}}
	//db
	_, err := colN.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
ListoUsuarios lista todos los usuarios registrados en el sistema,
*/
func ListoNodemcus(ID string, page int64, busqueda string) ([]*models.Nodemcu, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var resultados []*models.Nodemcu
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)
	qury := bson.M{
		//?i no se fija si son mayusculas y minusculas
		"dispositivo": bson.M{"$regex": `(?i)` + busqueda},
	}
	cursor, err := colN.Find(ctx, qury, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	for cursor.Next(ctx) {
		var s models.Nodemcu
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

func BuscoNodemcu(ID string) (models.Nodemcu, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	var nodemcu models.Nodemcu
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}
	err := colN.FindOne(ctx, condicion).Decode(&nodemcu)
	if err != nil {
		fmt.Println("registro no encontrado " + err.Error())
		return nodemcu, err
	}
	return nodemcu, nil
}

//RevisoSiExisteNodemcu recibe el email de parametro y revisa si ya existe en la db
func RevisoSiExisteNodemcu(tok string) (models.Nodemcu, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	condicion := bson.M{"token":tok,"status":"ok"}
	var resultado models.Nodemcu
	err := colN.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.Token
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
