package db

import (
	"context"
	"time"

	"github.com/MikeChain/registro-gastos-back/models"
	"go.mongodb.org/mongo-driver/bson"
)

// RevisarUsuario busca un usuario en la base de datos según el email
func RevisarUsuario(email string) (models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConexion.Database("movimientos")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	return resultado, err == nil
}
