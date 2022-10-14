package models

type Student struct {
	ID      uint   `json:"id" gorm:"index"`
	Name    string `json:"name" gorm:"type:varchar(255)"`
	Surname string `json:"surname" gorm:"type:varchar(255)"`
	Age     int8   `json:"age"`
}
