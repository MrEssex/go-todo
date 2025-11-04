package controllers

import (
	"context"
	"github.com/keystonedb/sdk-go/keystone"
	"github.com/keystonedb/sdk-go/proto"
	"github.com/mressex/go-todo/database"
	"github.com/mressex/go-todo/models"
	"log"
)

func GetAllTodos() ([]models.Todo, error) {
	todo := &models.Todo{}
	var todos []models.Todo
	actor := database.Actor()

	results, err := actor.List(context.Background(), keystone.Type(todo), []string{"title", "details", "completed"}, keystone.Limit(100, 0))
	if err != nil {
		return todos, err
	}

	err = keystone.UnmarshalToSlice(&todos, results...)
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func CreateTodo(todo string, details string) error {
	return database.Actor().Mutate(context.Background(), &models.Todo{Title: todo, Details: details, Completed: false})
}

func MarkTodoComplete(id string) error {
	todo := &models.Todo{Completed: true}
	todo.SetKeystoneID(id)

	return database.Actor().Mutate(context.Background(), todo, keystone.MutateProperties("completed"))
}

func MarkTodoIncomplete(id string) error {

	todo := &models.Todo{Completed: false}
	todo.SetKeystoneID(id)

	return database.Actor().Mutate(context.Background(), todo, keystone.MutateProperties("completed"))
}

func GetTodoByID(id string) (models.Todo, error) {
	todo := &models.Todo{}
	err := database.Actor().GetByID(context.Background(), id, todo, keystone.WithProperties())
	return *todo, err
}

func DeleteTodoByID(id string) error {
	x, err := database.Conn().Mutate(context.Background(), &proto.MutateRequest{
		Authorization: database.Actor().Authorization(),
		EntityId:      id,
		Mutation:      &proto.Mutation{State: proto.EntityState_Archived}})
	log.Println(x)
	return err
}
