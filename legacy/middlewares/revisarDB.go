package middlewares

import (
	"net/http"

	"github.com/MikeChain/registro-gastos-back/db"
)

// RevisarConexion es un middleware que verifica que la conexión con la db exista
func RevisarConexion(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.ConexionEstablecida() {
			http.Error(w, "Conexión fallida", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
