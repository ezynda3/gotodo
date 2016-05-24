package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/ezynda3/gotodo/models"
	"github.com/gin-gonic/gin"
)

// GetTasks endpoint
func GetTasks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTask endpoint
func PutTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		c.BindJSON(&task)
		id, err := models.PutTask(db, task.Name)
		if err == nil {
			c.JSON(http.StatusCreated, gin.H{
				"created": id,
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"error": "Could not create task",
			})
		}
	}
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteTask(db, id)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"deleted": id,
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"error": "Could not delete task",
			})
		}
	}
}
