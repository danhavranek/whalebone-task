package repositories

import (
	"github.com/danhavranek/whalebone-task/database"
	"github.com/danhavranek/whalebone-task/models"
)

func CreatePerson(person *models.Person) error {
	return database.DB.Create(person).Error
}

func GetPersonById(id string) (*models.Person, error) {
	var person models.Person
	err := database.DB.First(&person, id).Error
	if err != nil {
		return nil, err
	}
	return &person, nil
}
