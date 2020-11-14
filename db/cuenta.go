package db

import (
	"context"
	"log"
	"time"

	"github.com/MikeChain/registro-gastos-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertarCuenta guarda una cuenta en la DB
func InsertarCuenta(t models.Cuentas) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConexion.Database("movimientos")
	col := db.Collection("cuentas")

	registro := bson.M{
		"userid": t.UserID,
		"cuenta": t.Cuenta,
	}

	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}

// BuscarCuentas busca las cuentas que un usuario tiene registrado en la DB
func BuscarCuentas(ID string) ([]*models.Cuentas, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConexion.Database("movimientos")
	col := db.Collection("cuentas")

	var resultados []*models.Cuentas
	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetSort(bson.D{{Key: "cuenta", Value: 1}})

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.Cuentas
		err := cursor.Decode(&registro)

		if err != nil {
			return resultados, false
		}

		resultados = append(resultados, &registro)
	}

	return resultados, true
}
