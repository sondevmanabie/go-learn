package routes

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Get(c *gin.Context)
	Add(c *gin.Context)
	// Delete(c *gin.Context)
}

type Router interface {
	Init(r *gin.Engine)
}
