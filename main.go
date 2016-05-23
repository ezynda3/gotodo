package main

import (
	"database/sql"
	"net/http"

	"github.com/ezynda3/gotodo/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := initDB("storage.db")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.StaticFile("/", "./public/index.html")
	r.GET("/tasks", handlers.GetTasks(db))
	r.POST("/tasks", handlers.PutTask(db))
	r.DELETE("/tasks/:id", func(c *gin.Context) {})

	r.Run("localhost:8000")
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}
