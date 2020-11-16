package db

import (
	"context"
	"log"
	"time"

	"github.com/MikeChain/registro-gastos-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertarCuenta guarda una cuenta en la DB
func InsertarCuenta(t models.Cuentas) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	registro := bson.M{
		"userid": t.UserID,
		"cuenta": t.Cuenta,
	}

	result, err := cuentasCol.InsertOne(ctx, registro)
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

	var resultados []*models.Cuentas
	condicion := bson.M{
		"userid": ID,
	}

	cursor, err := cuentasCol.Find(ctx, condicion)

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

// ResultadosCuenta agrupa los movimientos de una cuenta de acuerdo a su tipo
