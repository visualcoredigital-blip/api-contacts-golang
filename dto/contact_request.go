package dto

type CreateContactRequest struct {
	Nombre      string          `json:"nombre" binding:"required"`
	Email       string          `json:"email" binding:"required,email"`
	Telefono    TelefonoRequest `json:"telefono" binding:"required"`
	Empresa     string          `json:"empresa"`
	Descripcion string          `json:"descripcion"`
	Estado      string          `json:"estado" binding:"required,oneof=nuevo contactado cerrado"`
}

type UpdateContactRequest struct {
	Nombre      *string          `json:"nombre"`
	Email       *string          `json:"email" binding:"omitempty,email"`
	Telefono    *TelefonoRequest `json:"telefono"`
	Empresa     *string          `json:"empresa"`
	Descripcion *string          `json:"descripcion"`
	Estado      *string          `json:"estado" binding:"omitempty,oneof=nuevo contactado cerrado"`
}

type TelefonoRequest struct {
	CodigoPais string `json:"codigoPais" binding:"required"`
	Numero     string `json:"numero" binding:"required"`
	Formateado string `json:"formateado"`
}