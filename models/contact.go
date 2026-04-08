package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre      string             `bson:"nombre" json:"nombre"`
	Email       string             `bson:"email" json:"email"`
	Telefono    Telefono           `bson:"telefono" json:"telefono"`
	Empresa     string             `bson:"empresa" json:"empresa"`
	Descripcion string             `bson:"descripcion" json:"descripcion"`
	Fecha       time.Time          `bson:"fecha" json:"fecha"`
	Estado      string             `bson:"estado" json:"estado"`
}

type Telefono struct {
	CodigoPais string `bson:"codigoPais" json:"codigoPais"`
	Formateado string `bson:"formateado" json:"formateado"`
	Numero     string `bson:"numero" json:"numero"`
}