package controller

import (
	"github.com/dzgntutar/ccgin/models"
	"github.com/dzgntutar/ccgin/service"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	Service service.IService[models.Student]
}

func (c StudentController) GetAll(context *gin.Context) {
	if err, studentList := c.Service.GetAll(); err != nil {
		context.JSON(400, err)
	} else {
		context.JSON(200, studentList)
	}
}

func (c StudentController) Create(context *gin.Context) {
	student := models.Student{}

	if err := context.ShouldBindJSON(&student); err != nil {
		context.AbortWithStatus(400)
	}

	if err := c.Service.Create(student); err != nil {
		context.AbortWithStatus(400)
	}
	context.Status(200)
}
