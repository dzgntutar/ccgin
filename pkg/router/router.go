package router

import (
	"github.com/dzgntutar/ccgin/controller"
	"github.com/dzgntutar/ccgin/repository"
	"github.com/dzgntutar/ccgin/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {
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
		student.POST("/", studentC.Create)
		student.GET("/:id", studentC.GetById)
	}

	return r
}
