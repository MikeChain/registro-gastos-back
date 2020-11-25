package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MikeChain/registro-gastos-back/db"
	"github.com/MikeChain/registro-gastos-back/models"
)

// InsertarSubrubro permite agregar un subrubro de un rubro
func InsertarSubrubro(w http.ResponseWriter, r *http.Request) {
	var subrubro struct {
		RubroID  string `bson:"rubroid" json:"rubroid"`
		Subrubro string `bson:"subrubro" json:"subrubro"`
	}

	err := json.NewDecoder(r.Body).Decode(&subrubro)

	rubros, _ := db.BuscarSubrubro(subrubro.RubroID)

	for _, a := range rubros {
		if a.Sububro == subrubro.Subrubro {
			http.Error(w, "El subrubro ya existe", 400)
			return
		}
	}

	if subrubro.Subrubro == "" {
		http.Error(w, "Envió un nombre de cuenta vacío", 400)
		return
	}

	registro := models.Subrubros{
		RubroID: subrubro.RubroID,
		Sububro: subrubro.Subrubro,
	}

	_, status, err := db.InsertarSubrubro(registro)
	if err != nil {
		http.Error(w, "Problemas al insertar el subrubro "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se pudo guardar el subrubro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// BuscarSubrubros permite obtener los subrubros que hay registrados en un rubro
func BuscarSubrubros(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("subrubro")

	if len(id) < 1 {
		http.Error(w, "Debe enviar el ID de rubro", http.StatusBadRequest)
	}

	respuesta, ok := db.BuscarSubrubro(id)
	if !ok {
		http.Error(w, "Error al buscar subrubros", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
