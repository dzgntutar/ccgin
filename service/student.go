package service

import "github.com/dzgntutar/ccgin/models"

type IStudentService interface {
	GetAll() ([]models.Student, error)
	Insert(student models.Student) error
	GetById(id int32) (models.Student, error)
}

type StudentService struct {
}

func GetAll() {

}

func Insert(student models.Student) {

}

func GetById(id int32) {

}
