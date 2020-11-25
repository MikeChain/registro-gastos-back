package routers

import (
	"encoding/json"
	"net/http"

	"github.com/MikeChain/registro-gastos-back/db"
	"github.com/MikeChain/registro-gastos-back/models"
)

// Registro se utiliza para agregar un nuevo usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El emial no puede estar vacío", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña es muy corta", 400)
		return
	}

	_, encontrado := db.RevisarUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario con ese email", 400)
		return
	}

	_, status, err := db.RegistroUsuario(t)
	if err != nil || !status {
		http.Error(w, "Error al guardar", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
