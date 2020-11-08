package db

import (
	"github.com/MikeChain/registro-gastos-back/models"
	"golang.org/x/crypto/bcrypt"
)

// Login revisa los datos en la DB
func Login(email, pass string) (models.Usuario, bool) {
	usuario, encontrado := RevisarUsuario(email)

	if !encontrado {
		return usuario, false
	}

	passBytes := []byte(pass)
	passDB := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passDB, passBytes)

	return usuario, err == nil
}
