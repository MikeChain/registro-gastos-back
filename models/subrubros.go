package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Subrubros es el modelo para las categorias en las que se clasifica un movimiento
type Subrubros struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RubroID string             `bson:"rubroid" json:"rubroid"`
	Sububro string             `bson:"subrubro" json:"subrubro"`
}
