package models

// ResultadosCuenta es el modelo en el que se muestra el resultado de una cuenta
type ResultadosCuenta struct {
	Ingresos  Dinero `bson:"ingresos,omitempty" json:"ingresos"`
	Egresos   Dinero `bson:"egresos,omitempty" json:"egresos"`
	Resultado Dinero `bson:"resultado,omitempty" json:"resultado"`
}
