package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/MikeChain/registro-gastos-back/db"
	"github.com/MikeChain/registro-gastos-back/models"
)

// InsertarMovimiento sirve para agregar un movimiento
func InsertarMovimiento(w http.ResponseWriter, r *http.Request) {
	var movimiento struct {
		CuentaID   string    `bson:"cuentaid" json:"cutnaid"`
		TipoID     string    `bson:"tipoid" json:"tipoid"`
		RubroID    string    `bson:"rubroid" json:"rubroid"`
		SubrubroID string    `bson:"subrubroid" json:"subrubroid"`
		Momento    time.Time `bson:"momento" json:"momento"`
		Detalle    string    `bson:"detalle" json:"detalle"`
	}
	err := json.NewDecoder(r.Body).Decode(&movimiento)

	if movimiento.CuentaID == "" {
		http.Error(w, "Debe seleccionar una cuenta", 400)
		return
	}

	if movimiento.TipoID == "" {
		http.Error(w, "Debe seleccionar un tipo de movimiento", 400)
		return
	}

	if movimiento.RubroID == "" {
		http.Error(w, "Debe seleccionar un rubro de movimiento", 400)
		return
	}

	var zero = time.Time{}
	if movimiento.Momento == zero {
		http.Error(w, "Debe establecer una fecha y hora del movimiento", 400)
		return
	}

	if movimiento.Detalle == "" {
		http.Error(w, "Debe especificar el detalle del movimiento", 400)
		return
	}

	registro := models.Movimientos{
		CuentaID:   movimiento.CuentaID,
		TipoID:     movimiento.TipoID,
		RubroID:    movimiento.RubroID,
		SubrubroID: movimiento.SubrubroID,
		Momento:    movimiento.Momento,
		Detalle:    movimiento.Detalle,
	}

	_, status, err := db.InsertarMovimiento(registro)
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

// BuscarMovimientos permite obtener los movimientos registrados de un usuario
func BuscarMovimientos(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el ID", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		pagina = 1
	}

	limit, erro := strconv.Atoi(r.URL.Query().Get("pagina"))
	if erro != nil {
		limit = 10
	}

	respuesta, ok := db.BuscarMovimientos(ID, pagina, limit)

	if !ok {
		http.Error(w, "Error al buscar movimientos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
