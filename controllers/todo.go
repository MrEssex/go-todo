package controllers

import (
	"github.com/mressex/go-todo/models"
)

func GetAllTodos() ([]models.Todo, error) {
	//todo := &models.Todo{}
	//actor := database.Actor()

	//err := actor.Get(context.Background(), keystone.ByEntityID(todo), todo, keystone.WithProperties())

	return []models.Todo{
		{Id: 1, Title: "Task 1", Details: "Task 1 details", Completed: false},
		{Id: 2, Title: "Task 2", Details: "Task 2 details", Completed: true},
		{Id: 3, Title: "Task 3", Details: "Task 3 details", Completed: false},
	}, nil
}

func CreateTodo(todo string, details string) error {
	return nil
}
