package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Cuentas es el modelo para las cuentas en las que se pueden realizar operaciones
type Cuentas struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Cuenta string             `bson:"cuenta" json:"cuenta"`
}
