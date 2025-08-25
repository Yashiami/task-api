package main

import (
	"github.com/gin-gonic/gin"
	"task-api/api_cmd"
	"task-api/database"
)

func main() {
	database.InitDB()
	router := gin.Default()
	router.POST("/create", api_cmd.CreateItem)
	router.GET("/todos", api_cmd.GetItems)
	router.PUT("/todos/:id", api_cmd.UpdateItem)
	router.Run("localhost:8080")

}
