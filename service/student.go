package service

import (
	"github.com/dzgntutar/ccgin/models"
	"github.com/dzgntutar/ccgin/repository"
)

type StudentService struct {
	Repository repository.IRepository[models.Student]
}

func (s StudentService) GetAll() ([]models.Student, error) {
	return s.Repository.GetAll()
}

func (s StudentService) Create(student models.Student) error {
	return s.Repository.Create(student)
}

func (s StudentService) GetById(id string) (models.Student, error) {
	return s.Repository.GetById(id)
}

func (s StudentService) Delete(id string) error {
	return s.Repository.Delete(id)
}
