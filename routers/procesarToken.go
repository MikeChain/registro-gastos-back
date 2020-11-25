package routers

import (
	"errors"
	"os"

	"github.com/MikeChain/registro-gastos-back/db"
	"github.com/MikeChain/registro-gastos-back/models"
	"github.com/dgrijalva/jwt-go"
)

// Email es el valor del correo del usuario autenticado
var Email string

// IDUsuario es el id del usuario autenticado
var IDUsuario string

// ProcesarToken permite extraer los valores del token del usuario
func ProcesarToken(tk string) (*models.Claim, bool, string, error) {
	clave := os.Getenv("KEY_JWT")
	if clave == "" {
		clave = "MiSuperDuperClaveSecretaParaGenerarMiTokenDeAutenticacion"
	}

	miClave := []byte(clave)
	claims := &models.Claim{}

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err != nil {
		return claims, false, "", err
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("token inválido")
	}

	_, encontrado := db.RevisarUsuario(claims.Email)

	if encontrado {
		Email = claims.Email
		IDUsuario = claims.ID.Hex()
	}

	return claims, encontrado, IDUsuario, nil
}
