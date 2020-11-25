package db

import (
	"context"
	"log"
	"time"

	"github.com/MikeChain/registro-gastos-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertarRubro guarda un rubro en la DB
func InsertarRubro(t models.Rubros) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	registro := bson.M{
		"tipoid": t.TipoID,
		"rubro":  t.Rubro,
		"userid": t.UserID,
	}

	result, err := rubrosCol.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}

// BuscarRubros busca los rubros que le pertenecen a un tipo de movimiento según el usuario
func BuscarRubros(IDuser, IDtipo string) ([]*models.Rubros, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var resultados []*models.Rubros
	condicion := bson.M{
		"userid": IDuser,
		"tipoid": IDtipo,
	}

	cursor, err := rubrosCol.Find(ctx, condicion)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.Rubros
		err := cursor.Decode(&registro)

		if err != nil {
			return resultados, false
		}

		resultados = append(resultados, &registro)
	}
	return resultados, true
}
