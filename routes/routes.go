package routes

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/mressex/go-todo/controllers"
	"github.com/mressex/go-todo/models"
	"github.com/mressex/go-todo/views"
	"log"
	"net/http"
)

func TodoRouter(r chi.Router) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		todos, err := controllers.GetAllTodos()
		if err != nil {
			todos = []models.Todo{}
		}

		templ.Handler(views.Home(todos)).ServeHTTP(w, r)
	})

	r.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "could not parse form", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		details := r.FormValue("details")
		err = controllers.CreateTodo(title, details)
		if err != nil {
			http.Error(w, "could not create todo", http.StatusInternalServerError)
			return
		}

		log.Println("Redirecting to /")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	r.Post("/todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		todo, err := controllers.GetTodoByID(id)
		if err != nil {
			http.Error(w, "could not get todo by id", http.StatusInternalServerError)
			return
		}

		if todo.Completed {
			err = controllers.MarkTodoIncomplete(id)
		} else {
			err = controllers.MarkTodoComplete(id)
		}

		if err != nil {
			http.Error(w, "could not mark todo complete", http.StatusInternalServerError)
			return
		}

		log.Println("Redirecting to /")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	r.Post("/todo/{id}/delete", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		err := controllers.DeleteTodoByID(id)
		if err != nil {
			http.Error(w, "could not delete todo by id", http.StatusInternalServerError)
			return
		}

		log.Println("Redirecting to /")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

}
