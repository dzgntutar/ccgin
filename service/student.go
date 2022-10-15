package service

import (
	"github.com/dzgntutar/ccgin/models"
	"github.com/dzgntutar/ccgin/repository"
)

type StudentService struct {
	Repository repository.IRepository[models.Student]
}

func (s StudentService) GetAll() ([]models.Student, error) {
	// fmt.Println()
	// fmt.Println("StudentService-GetAll -->", s.Repository)
	// fmt.Println()

	return s.Repository.GetAll()
}

func (s StudentService) Create(student models.Student) {
	s.Repository.Create(student)
}
