package service

import (
	"fmt"

	"github.com/dzgntutar/ccgin/repository"
)

type IStudentService interface {
	GetAll()
	// Insert(student models.Student) error
	// GetById(id int32) (models.Student, error)
}

type StudentService struct {
	StudentRepository repository.IStudentRepository
}

func (s StudentService) GetAll() {
	fmt.Println("StudentService --> GetAll")

	s.StudentRepository.GetAll()
}
