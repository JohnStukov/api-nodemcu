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

	var dbU = MongoCN.Database("nosql")
	var colU = dbU.Collection("usuarios")

/*InsertoRegistro es la parada final con la BD para insertar los datos del usuario */
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.Password, _ = EncriptarPassword(u.Password)
	result, err := colU.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

func BorroUsuario(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}

	_, err := colU.DeleteOne(ctx, condicion)
	return err
}

//ModificoRegistro permite modificar los datos del usuario
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	//armar el registro de la actualizacion de la db // Ejemplo 1
	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	if len(u.Imagen) > 0 {
		registro["imagen"] = u.Imagen
	}
	updateString := bson.M{
		"$set": registro,
	}
	//Proceso de actualizacion
	//convierte el ID de string a objID
	objID, _ := primitive.ObjectIDFromHex(ID)
	//filtar usuario
	filtro := bson.M{"_id": bson.M{"$eq": objID}}
	//db
	_, err := colU.UpdateOne(ctx, filtro, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
ListarUsuarios lista todos los usuarios registrados en el sistema,
*/
func ListarUsuarios(ID string, page int64, busqueda string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var resultados []*models.Usuario
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)
	qury := bson.M{
		//?i no se fija si son mayusculas y minusculas
		"nombre": bson.M{"$regex": `(?i)` + busqueda},
	}
	cursor, err := colU.Find(ctx, qury, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	for cursor.Next(ctx) {
		var s models.Usuario
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return resultados, false
		}
		s.Email = ""
		s.Password = ""
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

func BuscarUsuario(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	var usuario models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{
		"_id": objID,
	}
	err := colU.FindOne(ctx, condicion).Decode(&usuario)
	usuario.Password = ""
	if err != nil {
		fmt.Println("registro no encontrado " + err.Error())
		return usuario, err
	}
	return usuario, nil
}

//RevisarSiExisteUsuario recibe el email de parametro y revisa si ya existe en la db
func RevisarSiExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	condicion := bson.M{"email": email}
	var resultado models.Usuario
	err := colU.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}

