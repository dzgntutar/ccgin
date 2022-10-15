package service

import (
	"fmt"

	"github.com/dzgntutar/ccgin/repository"
)

type IStudentService interface {
	GetAll()
	Create()
}

type StudentService struct {
	Repository repository.IStudentRepository
}

func (s StudentService) GetAll() {
	fmt.Println()
	fmt.Println("StudentService-GetAll -->", s.Repository)
	fmt.Println()

	s.Repository.GetAll()
}

func (s StudentService) Create() {
	s.Repository.Create()
}
