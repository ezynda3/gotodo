package handlers

import (
	"database/sql"
	"net/http"

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
			c.JSON(http.StatusCreated, gin.H{
				"success": "Task created!",
			})
		}
	}
}
