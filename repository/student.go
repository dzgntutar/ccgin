package repository

import (
	"fmt"

	"github.com/dzgntutar/ccgin/models"
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
	fmt.Println()

	students := []models.Student{}

	query := r.Db.Select("students")

	if err := query.Take(&students).Error; err != nil {
		fmt.Println("Hata Repository-GetAll -->", err)
	}

	fmt.Println(students)
}
