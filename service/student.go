package service

import (
	"fmt"

	"github.com/dzgntutar/ccgin/repository"
)

type IStudentService interface {
	GetAll()
}

type StudentService struct {
	Repository repository.IStudentRepository
}

func (s StudentService) GetAll() {
	fmt.Println("StudentService-GetAll -->", s.Repository)
	fmt.Println()
	s.Repository.GetAll()
}
