package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string `json:"name" gorm:"type:varchar(255)"`
	Surname string `json:"surname" gorm:"type:varchar(255)"`
	Age     int32  `json:"age"`
}
