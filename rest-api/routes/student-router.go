package routes

import (
	"rest-api/controller"
	"rest-api/file_reader"

	"github.com/gin-gonic/gin"
)

type StudentRouter struct{}

func (router *StudentRouter) Init(r *gin.Engine) {
	c := &controller.StudentController{
		FileReader: &file_reader.EntityReader{
			Entity: "student",
			Path:   "data",
		},
	}
	r.GET("/student", c.Get)
	r.POST("/student", c.Add)
	// r.POST("student/:id", c.Delete)
}
