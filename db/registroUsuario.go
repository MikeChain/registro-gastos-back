package db

import (
	"context"
	"time"

	"github.com/MikeChain/registro-gastos-back/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// RegistroUsuario inserta un usuario en la base
func RegistroUsuario(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConexion.Database("movimientos")
	col := db.Collection("usuarios")

	u.Password, _ = encriptar(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}

func encriptar(texto string) (string, error) {
	costo := 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(texto), costo)

	return string(bytes), err
}
