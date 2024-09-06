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
			log.Println(err)
			todos = []models.Todo{}
		}

		templ.Handler(views.Home(todos)).ServeHTTP(w, r)
	})

	r.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "could not parse form", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		details := r.FormValue("details")
		err = controllers.CreateTodo(title, details)
		if err != nil {
			log.Println(err)
			http.Error(w, "could not create todo", http.StatusInternalServerError)
			return
		}

		log.Println("Redirecting to /")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

}
