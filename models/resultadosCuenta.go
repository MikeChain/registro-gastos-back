package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ResultadosCuenta es el modelo en el que se muestra el resultado de una cuenta
type ResultadosCuenta struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Ingresos  Dinero             `bson:"ingresos,omitempty" json:"ingresos"`
	Egresos   Dinero             `bson:"egresos,omitempty" json:"egresos"`
	Resultado Dinero             `bson:"resultado,omitempty" json:"resultado"`
}
