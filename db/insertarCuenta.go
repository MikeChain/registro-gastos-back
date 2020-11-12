package db

import (
	"context"
	"time"

	"github.com/MikeChain/registro-gastos-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
