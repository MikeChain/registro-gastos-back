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
		Cuenta string `bson:"cuenta" json:"cuenta"`
	}
	err := json.NewDecoder(r.Body).Decode(&cuenta)

	registro := models.Cuentas{
		UserID: IDUsuario,
		Cuenta: cuenta.Cuenta,
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

// BuscarCuentas permite obtener las cuentas registradas de un usuario
func BuscarCuentas(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el ID", http.StatusBadRequest)
		return
	}

	respuesta, ok := db.BuscarCuentas(ID)
	if !ok {
		http.Error(w, "Error al buscar cuentas", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
