package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConexion almacena el cliente de conexión a la DB
var MongoConexion = ConectarDB()
var clientOptions = options.Client().ApplyURI("")

// ConectarDB es la función para establecer la conexión con la base de datos
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión establecida")

	return client
}

// ConexionEstablecida revisa y establece por primera vez la conexión
func ConexionEstablecida() bool {
	err := MongoConexion.Ping(context.TODO(), nil)

	return err == nil
}
