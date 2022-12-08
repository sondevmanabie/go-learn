package main

import (
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	studentRouter := &routes.StudentRouter{}
	bookRouter := &routes.BookRouter{}
	studentRouter.Init(router)
	bookRouter.Init(router)
	router.Run("localhost:9090")
}
