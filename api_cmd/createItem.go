package api_cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"task-api/auth"
	"task-api/dataModel"
	"task-api/database"
)

func CreateItem(c *gin.Context) {
	userId, err := auth.GetUserId(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		log.Fatal(err)
		return
	}
	var newTodo dataModel.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		log.Fatal(err)
		return
	}
	_, err = database.DB.Exec("INSERT INTO todo (title, description,user_id) values ($1,$2,$3)", newTodo.Title, newTodo.Description, userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, newTodo)
		log.Fatal(err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newTodo)
}
