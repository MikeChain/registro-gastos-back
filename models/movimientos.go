package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Movimientos es el modelo para las categorias en las que se clasifica un movimiento
type Movimientos struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CuentaID   string             `bson:"cutnaid" json:"cutnaid"`
	TipoID     string             `bson:"tipoid" json:"tipoid"`
	RubroID    string             `bson:"rubroid" json:"rubroid"`
	SubrubroID string             `bson:"subrubroid" json:"subrubroid"`
}
