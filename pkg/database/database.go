package database

import (
	"fmt"

	"github.com/dzgntutar/ccgin/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() error {
	db, err := gorm.Open(sqlite.Open("ugin.db"))

	if err != nil {
		fmt.Println(err)
		return err
	}

	db.AutoMigrate(&models.Student{})

	return nil
}
