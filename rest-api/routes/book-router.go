package routes

import (
	"rest-api/controller"
	"rest-api/file_reader"

	"github.com/gin-gonic/gin"
)

type BookRouter struct{}

func (router *BookRouter) Init(r *gin.Engine) {
	bookCtl := &controller.BookController{
		FileReader: &file_reader.EntityReader{
			Entity: "book",
			Path:   "data",
		},
	}
	r.GET("/book", bookCtl.Get)
	r.POST("/book", bookCtl.Add)
}
