package handlers

import (
	"net/http"

	"api-contacts-golang/dto"
	"api-contacts-golang/models"
	"api-contacts-golang/repository"

	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	contacts, err := repository.GetAllContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, contacts)
}

func CreateContact(c *gin.Context) {
	var req dto.CreateContactRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	contact := models.Contact{
		Nombre: req.Nombre,
		Email:  req.Email,
		Telefono: models.Telefono{
			CodigoPais: req.Telefono.CodigoPais,
			Numero:     req.Telefono.Numero,
			Formateado: req.Telefono.Formateado,
		},
		Empresa:     req.Empresa,
		Descripcion: req.Descripcion,
		Estado:      req.Estado,
	}

	result, err := repository.CreateContact(contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetContactByID(c *gin.Context) {
	id := c.Param("id")

	contact, err := repository.GetContactByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "contacto no encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func UpdateContact(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateContactRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updated, err := repository.UpdateContactPartial(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteContact(c *gin.Context) {
	id := c.Param("id")

	err := repository.DeleteContact(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "contacto eliminado",
	})
}