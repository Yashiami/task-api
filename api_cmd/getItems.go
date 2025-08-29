package api_cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"task-api/auth"
	"task-api/dataModel"
	"task-api/database"
)

func GetItems(c *gin.Context) {
	userId, err := auth.GetUserId(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		log.Fatal(err)
		return
	}
	var allItems []dataModel.Todo
	rows, err := database.DB.Query("SELECT title,description FROM todo WHERE user_id = $1", userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Fatal(err)
		return
	}
	for rows.Next() {
		var currentTodo dataModel.Todo
		err = rows.Scan(&currentTodo.Title, &currentTodo.Description)
		if err != nil {
			log.Fatal(err)
		}
		allItems = append(allItems, currentTodo)
	}
	c.IndentedJSON(http.StatusOK, allItems)
}
