package models

import "errors"

// Dinero es el tipo de dato para manejar dinero
type Dinero struct {
	cantidad int64
	moneda   string
}

// Diferencia permite obtener la diferencia de dinero siempre que sea la misma moneda
func (d Dinero) Diferencia(n Dinero) (Dinero, error) {
	if d.moneda == n.moneda {
		x := d.cantidad - n.cantidad
		return Dinero{
			cantidad: x,
			moneda:   d.moneda,
		}, nil
	}
	return Dinero{}, errors.New("debe utilizar la misma moneda")
}

// Suma permite sumar dos cantidades de la misma moneda
func (d Dinero) Suma(b Dinero) (Dinero, error) {
	if d.moneda == b.moneda {
		x := d.cantidad + b.cantidad
		return Dinero{
			cantidad: x,
			moneda:   d.moneda,
		}, nil
	}
	return Dinero{}, errors.New("debe utilizar la misma moneda")
}

// EsMenorQue permite saber si una cantidad de dinero es menor a un entero
func (d Dinero) EsMenorQue(i int) bool {
	return d.cantidad < int64(1)
}
