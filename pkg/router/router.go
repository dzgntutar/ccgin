package router

import (
	"fmt"

	"github.com/dzgntutar/ccgin/controller"
	"github.com/dzgntutar/ccgin/repository"
	"github.com/dzgntutar/ccgin/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {

	fmt.Println("router-Init -->", db)
	fmt.Println()

	r := gin.Default()

	studentR := repository.StudentRepository{
		Db: db,
	}
	studentS := service.StudentService{
		Repository: studentR,
	}
	studentC := controller.StudentController{
		Service: studentS,
	}

	student := r.Group("/student")
	{
		student.GET("/", studentC.GetAll)
	}

	return r
}
