package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/tasks", func(c *gin.Context) {})
	r.POST("/tasks", func(c *gin.Context) {})
	r.DELETE("/tasks/:id", func(c *gin.Context) {})

	r.Run("localhost:8000")
}
