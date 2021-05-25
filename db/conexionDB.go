package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//MongoCN es el objeto de conexion a la db
var MongoCN = ConectarDB()
var clientOption = options.Client().ApplyURI("mongodb://root:123456@172.17.0.3:27017")

//ConetarDB ES LA FUNCION Q ME PERMITE CONECTAR A LA BASE DE DATOS
func ConectarDB() *mongo.Client {
	//CONECTA LA BASE DE DATOS, contexto es un espacio en memoria, q tiene una ejecucion, evita colgar la API, es un entorno de ejecucion
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la DB")
	return client
}

//RevisarConexion es el ping a la db
func RevisarConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
