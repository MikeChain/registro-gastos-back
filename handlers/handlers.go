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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
