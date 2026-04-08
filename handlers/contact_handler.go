package handlers

import (
	"net/http"

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
	var contact models.Contact

	// Bind JSON
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "JSON inválido",
		})
		return
	}

	// Validación mínima (no lo dejes pasar vacío)
	if contact.Nombre == "" || contact.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "nombre y email son obligatorios",
		})
		return
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
	var contact models.Contact

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "JSON inválido",
		})
		return
	}

	updated, err := repository.UpdateContact(id, contact)
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