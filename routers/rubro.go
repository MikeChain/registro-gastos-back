package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MikeChain/registro-gastos-back/db"
	"github.com/MikeChain/registro-gastos-back/models"
)

// InsertarRubro permite agregar un rubro de movimientos
func InsertarRubro(w http.ResponseWriter, r *http.Request) {
	var rubro struct {
		TipoID string `bson:"tipoid" json:"tipoid"`
		Rubro  string `bson:"rubro" json:"rubro"`
	}
	err := json.NewDecoder(r.Body).Decode(&rubro)

	rubros, _ := db.BuscarRubros(IDUsuario, rubro.TipoID)

	for _, a := range rubros {
		if a.Rubro == rubro.Rubro {
			http.Error(w, "El rubro ya existe", 400)
			return
		}
	}

	if rubro.Rubro == "" {
		http.Error(w, "Envió un nombre de cuenta vacío", 400)
		return
	}

	registro := models.Rubros{
		UserID: IDUsuario,
		Rubro:  rubro.Rubro,
		TipoID: rubro.TipoID,
	}

	_, status, err := db.InsertarRubro(registro)
	if err != nil {
		http.Error(w, "Problemas al insertar el rubro "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se pudo guardar el rubro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// BuscarRubros permite obtener los rubros registrados para un tipo de movimiento del usuario
func BuscarRubros(w http.ResponseWriter, r *http.Request) {
	IDrubro := r.URL.Query().Get("rubro")
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el ID", http.StatusBadRequest)
		return
	}

	if len(IDrubro) < 1 {
		http.Error(w, "Debe enviar el rubro", http.StatusBadRequest)
		return
	}

	respuesta, ok := db.BuscarRubros(ID, IDrubro)
	if !ok {
		http.Error(w, "Error al buscar rubros", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
