package routers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/MikeChain/registro-gastos-back/db"
	"github.com/MikeChain/registro-gastos-back/models"
	"github.com/dgrijalva/jwt-go"
)

// Login realiza los pasos para autenticar al usuario
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos inválidos: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email está vacío", 400)
		return
	}

	documento, existe := db.Login(t.Email, t.Password)
	if !existe {
		http.Error(w, "Datos invalidos", 400)
		return
	}

	jwtKey, err := generarJwt(documento)
	if err != nil {
		http.Error(w, "Error al generar token: "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

func generarJwt(t models.Usuario) (string, error) {
	clave := os.Getenv("KEY_JWT")
	if clave == "" {
		clave = "MiSuperDuperClaveSecretaParaGenerarMiTokenDeAutenticacion"
	}

	miClave := []byte(clave)

	payload := jwt.MapClaims{
		"email": t.Email,
		"id":    t.ID.Hex(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)

	return tokenStr, err
}
