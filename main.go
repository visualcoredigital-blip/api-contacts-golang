package main

import (
	"api-contacts-golang/config"
	"api-contacts-golang/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	// 🔴 ESTO ES OBLIGATORIO
	config.ConnectDB()

	router := gin.Default()

	router.GET("/contacts", handlers.GetContacts)
	router.POST("/contacts", handlers.CreateContact)
	router.GET("/contacts/:id", handlers.GetContactByID)
	router.PUT("/contacts/:id", handlers.UpdateContact)
	router.DELETE("/contacts/:id", handlers.DeleteContact)
	router.Run(":8080")
}