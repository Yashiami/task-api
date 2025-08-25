package api_cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"task-api/dataModel"
	"task-api/database"
)

func CreateItem(c *gin.Context) {
	var newTodo dataModel.Todo

	if err := c.BindJSON(newTodo); err != nil {
		return
	}
	stmt, err := database.DB.Prepare("INSERT INTO todo (title, description,user_id) values ($1,$2,$3)")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, newTodo)
		log.Fatal(err)
		return
	}
	if _, err := stmt.Exec(newTodo.Title, newTodo.Description, 1); err != nil {
		c.IndentedJSON(http.StatusBadRequest, newTodo)
		log.Fatal(err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newTodo)
}
