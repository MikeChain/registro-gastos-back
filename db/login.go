package db

import (
	"context"
	"time"

	"github.com/MikeChain/registro-gastos-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login revisa los datos en la DB
func Login(email, pass string) (models.Usuario, bool) {
	usuario, encontrado := revisarUsuario(email)

	if !encontrado {
		return usuario, false
	}

	passBytes := []byte(pass)
	passDB := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passDB, passBytes)

	return usuario, err == nil
}

func revisarUsuario(email string) (models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConexion.Database("movimientos")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	return resultado, err == nil
}
