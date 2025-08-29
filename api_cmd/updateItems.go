package api_cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"task-api/auth"
	"task-api/database"
)

func UpdateItem(c *gin.Context) {
	userId, err := auth.GetUserId(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		log.Fatal(err)
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Fatal(err)
		return
	}
	todo, err := findItemFromId(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Fatal(err)
		return
	}
	if err = c.BindJSON(todo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Fatal(err)
		return
	}
	_, err = database.DB.Exec("UPDATE todo SET title = $1, description = $2 WHERE user_id = $3 AND id = $4", todo.Title, todo.Description, userId, id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Fatal(err)
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}
