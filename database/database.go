package database

import (
	"github.com/danhavranek/whalebone-task/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const dbFileName string = "test.db"

func Init() error {
	var err error

	DB, err = gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.Person{})
	if err != nil {
		return err
	}

	return nil
}
