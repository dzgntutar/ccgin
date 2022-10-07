package controller

import (
	"fmt"

	"github.com/dzgntutar/ccgin/service"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	StudentService service.IStudentService
}

func (controller StudentController) GetAll(c *gin.Context) {
	fmt.Println("StudentController --> GetAll")
	controller.StudentService.GetAll()
}

// func (controller StudentController) GetById(c *gin.Context) {
// 	controller.service.GetById(1)
// }

// func (controller StudentController) Create(c *gin.Context) {
// 	controller.Create(c)
// }
