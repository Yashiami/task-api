package api_cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"task-api/dataModel"
	"task-api/database"
)

func GetItems(c *gin.Context) {
	var allItems []dataModel.Todo
	rows, err := database.DB.Query("SELECT title,description FROM todo")
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
