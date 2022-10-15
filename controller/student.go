package controller

import (
	"fmt"

	"github.com/dzgntutar/ccgin/service"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	Service service.IStudentService
}

func (c StudentController) GetAll(context *gin.Context) {
	fmt.Println()
	fmt.Println("StudentController-GetAll -->", c.Service)
	fmt.Println()

	if err, studentList := c.Service.GetAll(); err != nil {
		context.JSON(400, err)
	} else {
		context.JSON(200, studentList)
	}

}

func (c StudentController) Create(context *gin.Context) {
	c.Service.Create()
}
