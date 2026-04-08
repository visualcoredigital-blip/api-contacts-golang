package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"api-contacts-golang/config"
	"api-contacts-golang/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllContacts() ([]models.Contact, error) {
	collection := config.DB.Collection("contacts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var contacts []models.Contact

	for cursor.Next(ctx) {
		var contact models.Contact
		if err := cursor.Decode(&contact); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func CreateContact(contact models.Contact) (models.Contact, error) {
	collection := config.DB.Collection("contacts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Generar ID y fecha
	contact.ID = primitive.NewObjectID()
	contact.Fecha = time.Now()

	_, err := collection.InsertOne(ctx, contact)
	if err != nil {
		return contact, err
	}

	return contact, nil
}

func GetContactByID(id string) (models.Contact, error) {
	collection := config.DB.Collection("contacts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Contact{}, err
	}

	var contact models.Contact
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&contact)
	if err != nil {
		return models.Contact{}, err
	}

	return contact, nil
}

func UpdateContact(id string, contact models.Contact) (models.Contact, error) {
	collection := config.DB.Collection("contacts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Contact{}, err
	}

	update := bson.M{
		"$set": bson.M{
			"nombre":      contact.Nombre,
			"email":       contact.Email,
			"telefono":    contact.Telefono,
			"empresa":     contact.Empresa,
			"descripcion": contact.Descripcion,
			"estado":      contact.Estado,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return models.Contact{}, err
	}

	// devolver actualizado
	return GetContactByID(id)
}

func DeleteContact(id string) error {
	collection := config.DB.Collection("contacts")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}