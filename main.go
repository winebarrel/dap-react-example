package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{
	{ID: 1, Task: "Learn Go"},
	{ID: 2, Task: "Build a Web App"},
}

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})
	router.GET("/api/todos", getTodos) // ToDo の一覧取得
	router.Run(":8080")
}
