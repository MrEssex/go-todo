package controllers

import (
	"github.com/mressex/go-todo/database"
	"github.com/mressex/go-todo/models"
)

func GetAllTodos() ([]models.Todo, error) {
	//todo := &models.Todo{}
	//actor := database.Actor()

	//err := actor.Get(context.Background(), keystone.ByEntityID(todo), todo, keystone.WithProperties())

	todos := []models.Todo{}
	statement := "SELECT id, title, details, completed FROM todos;"
	rows, err := database.DB.Query(statement)
	if err != nil {
		return todos, err
	}

	for rows.Next() {
		todo := models.Todo{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Details, &todo.Completed)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func CreateTodo(todo string, details string) error {
	statement := "INSERT INTO todos (title, details, completed) VALUES (?, ?, ?);"
	_, err := database.DB.Exec(statement, todo, details, false)
	if err != nil {
		return err
	}
	return nil
}

func MarkTodoComplete(id int) error {
	statement := "UPDATE todos SET completed = ? WHERE id = ?;"
	_, err := database.DB.Exec(statement, true, id)
	if err != nil {
		return err
	}
	return nil
}
