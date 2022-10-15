package repository

import (
	"fmt"

	"github.com/dzgntutar/ccgin/models"
	"gorm.io/gorm"
)

type IStudentRepository interface {
	GetAll()
	Create()
}

type StudentRepository struct {
	Db *gorm.DB
}

func (r StudentRepository) GetAll() {
	fmt.Println()
	fmt.Println("Student Repository --> GetAll")
	fmt.Println("Student Repository ::", r.Db)
	fmt.Println()

	students := []models.Student{}

	query := r.Db.Select("students.*")

	if err := query.Find(&students).Error; err != nil {
		fmt.Println("Hata Repository-GetAll -->", err)
	}

	fmt.Println(students)
}

func (r StudentRepository) Create() {
	student := models.Student{
		Name:    "Ali",
		Surname: "Kaya",
		Age:     35,
	}
	if err := r.Db.Save(&student).Error; err != nil {
		fmt.Println("Hata Repository-Create", err)
	}
}
