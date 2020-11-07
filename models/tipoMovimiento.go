package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// TipoMovimiento es el modelo para los tipos de movimiento realizables
type TipoMovimiento struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Cuenta string             `bson:"cuenta" json:"cuenta"`
}
