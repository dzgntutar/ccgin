package service

import (
	"fmt"

	"github.com/dzgntutar/ccgin/models"
	"github.com/dzgntutar/ccgin/repository"
)

type IStudentService interface {
	GetAll() (error, []models.Student)
	Create()
}

type StudentService struct {
	Repository repository.IStudentRepository
}

func (s StudentService) GetAll() (error, []models.Student) {
	fmt.Println()
	fmt.Println("StudentService-GetAll -->", s.Repository)
	fmt.Println()

	return s.Repository.GetAll()
}

func (s StudentService) Create() {
	s.Repository.Create()
}
