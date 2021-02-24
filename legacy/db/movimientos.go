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

// InsertarMovimiento guarda el movimiento en la DB
func InsertarMovimiento(t models.Movimientos) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	registro := bson.M{
		"cuentaid":   t.CuentaID,
		"tipoid":     t.TipoID,
		"rubroid":    t.RubroID,
		"subrubroid": t.SubrubroID,
		"momento":    t.Momento,
	}

	result, err := movimientosCol.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}

// BuscarMovimientos busca los movimientos que le pertenecen a un usuario
func BuscarMovimientos(ID string, pag int, limite int) ([]*models.Movimientos, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	limit := int64(limite)
	pagina := int64(pag)

	var resultados []*models.Movimientos
	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(limit)
	opciones.SetSort(bson.D{{Key: "momento", Value: -1}})
	opciones.SetSkip((pagina - 1) * limit)

	cursor, err := movimientosCol.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.Movimientos
		err := cursor.Decode(&registro)

		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}

	return resultados, true
}
