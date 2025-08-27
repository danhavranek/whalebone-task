package repositories

import (
	"github.com/danhavranek/whalebone-task/database"
	"github.com/danhavranek/whalebone-task/models"
)

func CreatePerson(person *models.Person) error {
	// TODO: keep original offset in timestamp
	return database.DB.Create(person).Error
}

func GetPersonById(externalId string) (*models.Person, error) {
	var person models.Person
	err := database.DB.Where("external_id = ?", externalId).First(&person).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}
