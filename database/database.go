package database

import (
	"github.com/danhavranek/whalebone-task/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const dbPath string = "app/data/app.db"

func Init() error {
	var err error

	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.Person{})
	if err != nil {
		return err
	}

	return nil
}
