package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// tipoMovimiento es el modelo para los tipos de movimiento realizables
type tipoMovimiento struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Tipo string             `bson:"tipo" json:"tipo"`
}

var idIngresos, supErr = primitive.ObjectIDFromHex("000000000000000000000001")
var idEgresos, _ = primitive.ObjectIDFromHex("000000000000000000000002")
var idTransferenciaN, _ = primitive.ObjectIDFromHex("000000000000000000000003")
var idTransferenciaP, _ = primitive.ObjectIDFromHex("000000000000000000000004")

// Ingresos indica entradas de dinero a una cuenta
var Ingresos tipoMovimiento = tipoMovimiento{
	idIngresos,
	"Ingresos",
}

// Egresos indica salidas de dinero a una cuenta
var Egresos tipoMovimiento = tipoMovimiento{
	idEgresos,
	"Egresos",
}

// TransferenciaN indica salidas de dinero a una cuenta que entran en otra
var TransferenciaN tipoMovimiento = tipoMovimiento{
	idTransferenciaN,
	"Transferencia (-)",
}

// TransferenciaP indica entradas de dinero a una cuenta que salen de otra
var TransferenciaP tipoMovimiento = tipoMovimiento{
	idTransferenciaP,
	"Transferencia (+)",
}
