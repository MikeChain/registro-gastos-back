package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConexion almacena el cliente de conexión a la DB
var MongoConexion = ConectarDB()
var url = os.Getenv("mongoConection")

// ConectarDB es la función para establecer la conexión con la base de datos
func ConectarDB() *mongo.Client {

	if url == "" {
		url = "mongodb://127.0.0.1:27017/?gssapiServiceName=mongodb"
	}

	clientOptions := options.Client().ApplyURI(url)

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

// database
var db = MongoConexion.Database("movimientos")
var cuentasCol = db.Collection("cuentas")
var rubrosCol = db.Collection("rubros")
var subrubrosCol = db.Collection("subrubros")
