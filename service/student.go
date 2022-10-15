package service

import (
	"fmt"

	"github.com/dzgntutar/ccgin/models"
	"github.com/dzgntutar/ccgin/repository"
)

type IStudentService interface {
	GetAll() ([]models.Student, error)
	Create()
}

type StudentService struct {
	Repository repository.IStudentRepository
}

func (s StudentService) GetAll() ([]models.Student, error) {
	fmt.Println()
	fmt.Println("StudentService-GetAll -->", s.Repository)
	fmt.Println()

	return s.Repository.GetAll()
}

func (s StudentService) Create() {
	s.Repository.Create()
}
