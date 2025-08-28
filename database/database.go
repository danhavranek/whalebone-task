package database

import (
	"os"
	"path/filepath"

	"github.com/danhavranek/whalebone-task/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DbPath string = "app/data/app.db"

func Init() error {
	// Create DB path if not exists
	err := os.MkdirAll(filepath.Dir(DbPath), 0755)
	if err != nil {
		return err
	}

	DB, err = gorm.Open(sqlite.Open(DbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.Person{})
	if err != nil {
		return err
	}

	return nil
}
