package router

import (
	"github.com/dzgntutar/ccgin/controller"
	"github.com/dzgntutar/ccgin/repository"
	"github.com/dzgntutar/ccgin/service"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	studentR := repository.StudentRepository{}
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
