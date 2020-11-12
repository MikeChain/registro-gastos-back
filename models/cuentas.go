package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Cuentas es el modelo para las cuentas en las que se pueden realizar operaciones
type Cuentas struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID string             `bson:"userid,omitempty" json:"userid"`
	Cuenta string             `bson:"cuenta" json:"cuenta"`
}
