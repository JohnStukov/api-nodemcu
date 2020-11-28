package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*Usuario es el modelo de usuario en mongodb*/
type Usuario struct {
	ID            primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Nombre        string             `bson:"nombre" json:"nombre, omitempty"`
	Apellidos     string             `bson:"apellidos" json:"apellidos, omitempty"`
	FechaRegistro time.Time          `bson:"fechaRegistro" json:"fechaRegistro, omitempty"`
	Email         string             `bson:"email" json:"email"`
	Password      string             `bson:"password" json:"password, omitempty"`
	Imagen        string             `bson:"imagen" json:"imagen, omitempty"`
}
