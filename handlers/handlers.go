package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/MikeChain/registro-gastos-back/middlewares"
	"github.com/MikeChain/registro-gastos-back/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores genera los enpoint
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/login", middlewares.RevisarConexion(routers.Login)).Methods("POST")
	router.HandleFunc("/registro", middlewares.RevisarConexion(routers.Registro)).Methods("POST")

	// endpoints para cuentas
	router.HandleFunc("/cuenta", middlewares.RevisarConexion(middlewares.ValidarJWT(routers.InsertarCuenta))).Methods("POST")
	router.HandleFunc("/cuenta", middlewares.RevisarConexion(middlewares.ValidarJWT(routers.BuscarCuentas))).Methods("GET")

	// endpoints para rubros
	router.HandleFunc("/rubro", middlewares.RevisarConexion(middlewares.ValidarJWT(routers.InsertarRubro))).Methods("POST")
	router.HandleFunc("/rubro", middlewares.RevisarConexion(middlewares.ValidarJWT(routers.BuscarRubros))).Methods("GET")

	// endpoints para subrubros
	router.HandleFunc("/subrubro", middlewares.RevisarConexion(middlewares.ValidarJWT(routers.InsertarSubrubro))).Methods("POST")
	router.HandleFunc("/subrubro", middlewares.RevisarConexion(middlewares.ValidarJWT(routers.BuscarSubrubros))).Methods("GET")

	// endpoints para movimientos
	router.HandleFunc("/movimientos", middlewares.RevisarConexion(middlewares.ValidarJWT(routers.InsertarMovimiento))).Methods("POST")
	router.HandleFunc("/movimientos", middlewares.RevisarConexion(middlewares.ValidarJWT(routers.BuscarMovimientos))).Methods("GET")

	// TODO: movimientos especiales:
	// resultados por cuenta
	// router.HandleFunc("/cuentas/movimientos}", middlewares.RevisarConexion(middlewares.ValidarJWT(routers.BuscarMovimientos))).Methods("GET")
	// movimientos por cuenta
	// router.HandleFunc("/movimientos/cuentas/{ID:[a-zA-Z0-9_]+}", middlewares.RevisarConexion(middlewares.ValidarJWT(routers.BuscarMovimientos))).Methods("GET")
	// movimientos por tipo
	// movimientos por fecha (YTD, MTD)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
