package controller

import (
	"net/http"
	"rest-api/ent"
	"rest-api/file_reader"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	FileReader file_reader.FileReader
}

func (c *StudentController) Get(ctx *gin.Context) {
	var res []ent.Student
	data, err := c.FileReader.Read(&res)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, data)
}

func (c *StudentController) Add(ctx *gin.Context) {
	var newStudent ent.Student
	if err := ctx.BindJSON(&newStudent); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid params"})
		return
	}
	var students []ent.Student
	_, err := c.FileReader.Read(&students)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
		return
	}
	students = append(students, newStudent)
	if err := c.FileReader.Write(&students); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
}

// func (c *StudentController) Delete(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	data, err := c.FileReader.Read()
// 	if err != nil {
// 		ctx.IndentedJSON(http.StatusNotFound, err)
// 		return
// 	}
// 	temp := *data
// 	var v, ok interface{} = temp.(type)
// 	for i, v := range *data {
// 		temp := *ent.Student(v)
// 		if temp == id {
// 			index = i
// 		}
// 	}
// }
