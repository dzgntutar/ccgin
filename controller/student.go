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
	fmt.Println("StudentController --> GetAll")
	c.Service.GetAll()
}

// func (controller StudentController) GetById(c *gin.Context) {
// 	controller.service.GetById(1)
// }

// func (controller StudentController) Create(c *gin.Context) {
// 	controller.Create(c)
// }
