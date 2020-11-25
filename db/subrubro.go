package db

import (
	"context"
	"log"
	"time"

	"github.com/MikeChain/registro-gastos-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertarSubrubro guarda el subrubro en la DB
func InsertarSubrubro(t models.Subrubros) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	registro := bson.M{
		"rubroid":  t.RubroID,
		"subrubro": t.Sububro,
	}

	result, err := subrubrosCol.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}

// BuscarSubrubro busca los subrubros que le pertenecen a un rubro
// El rubro ya pertenece al usuario, por lo que no es necesario usar al usuario para búsqueda
func BuscarSubrubro(id string) ([]*models.Subrubros, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var resultados []*models.Subrubros
	condicion := bson.M{
		"rubroid": id,
	}

	cursor, err := subrubrosCol.Find(ctx, condicion)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.Subrubros
		err := cursor.Decode(&registro)

		if err != nil {
			return resultados, false
		}

		resultados = append(resultados, &registro)
	}
	return resultados, true
}
