package controllers

import (
	"github.com/mressex/go-todo/database"
	"github.com/mressex/go-todo/models"
)

func GetAllTodos() ([]models.Todo, error) {
	//todos := &models.Todo{}
	//actor := database.Actor()
	//err := actor.Get(context.Background(), keystone.ByEntityID(todos, todos.GetKeystoneID()), todos, keystone.WithProperties())

	var todos []models.Todo
	statement := "SELECT id, title, details, completed FROM todos;"
	rows, err := database.DB.Query(statement)
	if err != nil {
		return todos, err
	}

	var (
		id        int
		title     string
		details   string
		completed bool
	)

	for rows.Next() {
		err = rows.Scan(&id, &title, &details, &completed)
		if err != nil {
			return todos, err
		}
		todo := models.Todo{
			ID:        id,
			Title:     title,
			Details:   details,
			Completed: completed,
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

func MarkTodoIncomplete(id int) error {
	statement := "UPDATE todos SET completed = ? WHERE id = ?;"
	_, err := database.DB.Exec(statement, false, id)
	if err != nil {
		return err
	}
	return nil
}

func GetTodoByID(id int) (models.Todo, error) {
	statement := "SELECT id, title, details, completed FROM todos WHERE id = ?;"
	row := database.DB.QueryRow(statement, id)

	var (
		title     string
		details   string
		completed bool
	)

	err := row.Scan(&id, &title, &details, &completed)
	if err != nil {
		return models.Todo{}, err
	}

	todo := models.Todo{
		ID:        id,
		Title:     title,
		Details:   details,
		Completed: completed,
	}

	return todo, nil
}
