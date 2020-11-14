package middlewares

import (
	"net/http"

	"github.com/MikeChain/registro-gastos-back/routers"
)

//ValidarJWT es un middleware para validar los JWT
func ValidarJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesarToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
