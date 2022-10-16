package models

import "gorm.io/gorm"

type Human struct {
	gorm.Model
	Name    string `json:"Name" gorm:"type:varchar(255)"`
	Surname string `json:"Surname" gorm:"type:varchar(255)"`
	Age     int32  `json:"Age"`
}

type Teacher struct {
	Human
	Branch string
	Salary float32
}

type Manager struct {
	Human
	Salary float32
}

type Student struct {
	Human
	//Lessons    []Lesson
	//Department Department
}

type Employee struct {
	Human
	Salary float32
}
