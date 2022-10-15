package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string `json:"Name" gorm:"type:varchar(255)"`
	Surname string `json:"Surname" gorm:"type:varchar(255)"`
	Age     int32  `json:"Age"`
}
