package main

import (
	"log"

	"github.com/MikeChain/registro-gastos-back/db"
	"github.com/MikeChain/registro-gastos-back/handlers"
)

func main() {
	if db.ConexionEstablecida() {
		handlers.Manejadores()
	} else {
		log.Fatal("Sin conexión con la base")
	}
}
