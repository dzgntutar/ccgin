package router

import (
	"github.com/dzgntutar/ccgin/controller"
	"github.com/dzgntutar/ccgin/repository"
	"github.com/dzgntutar/ccgin/service"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	sRepository := repository.StudentRepository{
		Db: "MyDb",
	}
	sService := service.StudentService{
		StudentRepository: sRepository,
	}
	sController := controller.StudentController{
		StudentService: sService,
	}

	student := r.Group("/student")
	{
		student.GET("/", sController.GetAll)
		//student.GET("/:id", controller.GetById)
		//student.POST("/", controller.Create)
	}

	return r
}
