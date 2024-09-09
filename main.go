package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mressex/go-todo/database"
	"github.com/mressex/go-todo/routes"
	"net/http"
)

func main() {
	database.InitKeyStone()
	defer database.CloseKeyStone()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Group(routes.TodoRouter)

	http.ListenAndServe(":3030", r)
}
