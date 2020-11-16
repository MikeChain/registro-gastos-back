package models

// Dinero es el tipo de dato para manejar dinero
type Dinero struct {
	cantidad int64
	moneda   string
}

// Multiplicar permite multiplicar la cantidad de dinero d por un entero f
func (d Dinero) Multiplicar(f int64) Dinero {
	x := d.cantidad * f
	return Dinero{
		cantidad: x,
		moneda:   d.moneda,
	}
}
