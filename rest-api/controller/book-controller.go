package controller

import (
	"net/http"
	"rest-api/ent"
	"rest-api/file_reader"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	FileReader file_reader.FileReader
}

func (c *BookController) Get(ctx *gin.Context) {
	var res []ent.Book
	_, err := c.FileReader.Read(&res)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, res)
}
func (c *BookController) Add(ctx *gin.Context) {
	var newBook ent.Book
	if err := ctx.BindJSON(&newBook); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid params"})
		return
	}
	var books []ent.Book
	_, err := c.FileReader.Read(&books)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
		return
	}
	books = append(books, newBook)
	if err := c.FileReader.Write(&books); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
}
