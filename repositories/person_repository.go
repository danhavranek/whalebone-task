package repositories

import (
	"github.com/danhavranek/whalebone-task/database"
	"github.com/danhavranek/whalebone-task/models"
	"github.com/google/uuid"
)

func CreatePerson(person *models.Person) error {
	return database.DB.Create(person).Error
}

func GetPersonById(externalId uuid.UUID) (*models.Person, error) {
	var person models.Person
	err := database.DB.Where("external_id = ?", externalId).First(&person).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}
