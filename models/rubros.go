package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Rubros es el modelo para las categorias en las que se clasifica un movimiento
type Rubros struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TipoID string             `bson:"tipoid" json:"tipoid"`
	Rubro  string             `bson:"rubro" json:"rubro"`
}
