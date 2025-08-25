package api_cmd

import (
	"errors"
	"task-api/dataModel"
	"task-api/database"
)

func findItemFromId(id int) (*dataModel.Todo, error) {
	row := database.DB.QueryRow(`SELECT id,title,description FROM todo WHERE user_id = 1`)
	var currentTodo dataModel.Todo
	var idTodo int
	err := row.Scan(&idTodo, &currentTodo.Title, &currentTodo.Description)
	if err != nil {
		return nil, err
	}
	if idTodo == id {
		return &currentTodo, nil
	}
	return nil, errors.New("cannot find todo item")
}
