package router

import (
	"github.com/dzgntutar/ccgin/controller"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	student := r.Group("/student")
	{
		student.GET("/", controller.GetAll)
		student.GET("/:id", controller.GetById)
		student.POST("/", controller.Create)
	}

	return r
}
