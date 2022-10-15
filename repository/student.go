package repository

import (
	"github.com/dzgntutar/ccgin/models"
	"gorm.io/gorm"
)

type StudentRepository struct {
	Db *gorm.DB
}

func (r StudentRepository) GetAll() ([]models.Student, error) {
	students := []models.Student{}

	query := r.Db.Select("students.*")

	if err := query.Find(&students).Error; err != nil {
		return students, err
	}

	return students, nil
}

func (r StudentRepository) Create(student models.Student) error {

	if err := r.Db.Save(&student).Error; err != nil {
		return err
	}
	return nil
}
