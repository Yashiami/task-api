package api_cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"task-api/database"
)

func DeleteItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		log.Fatal(err)
		return
	}
	//userId := findRequestUserId()
	userId := 2
	rowCheck := database.DB.QueryRow(`SELECT title FROM todo WHERE user_id = $1 AND id = $2`, userId, id)
	var titleTodo string
	if rowCheck.Scan(&titleTodo) != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Cannot find specified item"})
		return
	} else {
		if _, err = database.DB.Exec(`DELETE FROM todo WHERE id = $1`, id); err != nil {
			c.IndentedJSON(http.StatusBadRequest, nil)
			log.Fatal(err)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, nil)
}
