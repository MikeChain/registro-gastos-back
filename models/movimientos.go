package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Movimientos es el modelo para los registros de movimiento
type Movimientos struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CuentaID   string             `bson:"cuentaid" json:"cutnaid"`
	TipoID     string             `bson:"tipoid" json:"tipoid"`
	RubroID    string             `bson:"rubroid" json:"rubroid"`
	SubrubroID string             `bson:"subrubroid" json:"subrubroid"`
	Momento    time.Time          `bson:"momento" json:"momento"`
	Detalle    string             `bson:"detalle" json:"detalle"`
}
