package database

import (
	"fmt"

	"github.com/dzgntutar/ccgin/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() error {
	db, err := gorm.Open(sqlite.Open("myDatabase.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return err
	}

	db.AutoMigrate(&models.Student{})

	DB = db

	return nil
}
