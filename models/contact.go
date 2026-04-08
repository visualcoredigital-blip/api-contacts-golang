package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre      string             `bson:"nombre" json:"nombre" binding:"required"`
	Email       string             `bson:"email" json:"email" binding:"required,email"`
	Telefono    Telefono           `bson:"telefono" json:"telefono"`
	Empresa     string             `bson:"empresa" json:"empresa"`
	Descripcion string             `bson:"descripcion" json:"descripcion"`
	Fecha       time.Time          `bson:"fecha" json:"fecha"`
	Estado      string             `bson:"estado" json:"estado" binding:"required,oneof=Pendiente Contactado"`
}

type Telefono struct {
	CodigoPais string `bson:"codigoPais" json:"codigoPais"`
	Formateado string `bson:"formateado" json:"formateado"`
	Numero     string `bson:"numero" json:"numero"`
}

type UpdateContactRequest struct {
	Nombre      *string    `json:"nombre"`
	Email       *string    `json:"email"`
	Telefono    *Telefono  `json:"telefono"`
	Empresa     *string    `json:"empresa"`
	Descripcion *string    `json:"descripcion"`
	Estado      *string    `json:"estado"`
}