package repository

import "fmt"

type IStudentRepository interface {
	GetAll()
}

type StudentRepository struct {
	Db string
}

func (r StudentRepository) GetAll() {
	fmt.Println("Student Repository --> GetAll")

	fmt.Println("Student Repository ::", r.Db)
}
