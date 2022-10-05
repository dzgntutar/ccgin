package controller

import (
	"github.com/dzgntutar/ccgin/models"
	"github.com/dzgntutar/ccgin/service"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	studentService service.IStudentService
}

func GetAll(c *gin.Context) {

	std := models.Student{
		Name:    "Düzgün",
		Surname: "Tutar",
		Age:     35,
	}

	c.JSON(200, std)
}

func GetById(c *gin.Context) {

}

func Create(c *gin.Context) {

}
