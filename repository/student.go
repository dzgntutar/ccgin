package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type IStudentRepository interface {
	GetAll()
}

type StudentRepository struct {
	Db *gorm.DB
}

func (r StudentRepository) GetAll() {
	fmt.Println("Student Repository --> GetAll")

	fmt.Println("Student Repository ::", r.Db)
}
