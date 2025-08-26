package main

import (
	"github.com/gin-gonic/gin"
	"task-api/api_cmd"
	"task-api/auth"
	"task-api/database"
)

func main() {
	database.InitDB()
	router := gin.Default()
	defer database.DB.Close()
	router.POST("/add-todo", api_cmd.CreateItem)
	router.POST("/sighup", auth.CreateUser)
	router.POST("/auth", auth.AuthorizationUser)
	router.GET("/todo", api_cmd.GetItems)
	router.PUT("/todo/:id", api_cmd.UpdateItem)
	router.DELETE("todo/:id", api_cmd.DeleteItem)
	router.Run("localhost:8080")

}
