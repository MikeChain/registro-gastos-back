package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MikeChain/registro-gastos-back/db"
	"github.com/MikeChain/registro-gastos-back/models"
)

// InsertarCuenta permite agregar una cuenta para realizar movimientos
func InsertarCuenta(w http.ResponseWriter, r *http.Request) {
	var cuenta struct {
		Mensaje string `bson:"mensaje" json:"mensaje"`
	}
	err := json.NewDecoder(r.Body).Decode(&cuenta)

	registro := models.Cuentas{
		UserID: IDUsuario,
		Cuenta: cuenta.Mensaje,
	}

	_, status, err := db.InsertarCuenta(registro)
	if err != nil {
		http.Error(w, "Problemas al insertar la cuenta "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se pudo guardar la cuenta", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
