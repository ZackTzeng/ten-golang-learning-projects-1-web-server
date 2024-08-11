package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/ZackTzeng/ten-golang-learning-projects-1-web-server/internal/todo"
)

func NewRouter() *chi.Mux {
	ts := todo.NewTodoStore()
	firstTodo := todo.NewTodo("First Todo", "This is the first todo")
	firstTodo.PrintTodo()
	ts.AddTodo(*firstTodo)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome to Todo API!"))
		})
		r.Route("/todos", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode(ts.GetTodoList())
			})
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				// request := map[string]string{}
				// json.NewDecoder(r.Body).Decode(&request)
				// newTodo := todo.NewTodo(request["subject"], request["content"])
				// ts.AddTodo(*newTodo)
				// w.Write([]byte("Added new todo"))

				data := &TodoRequest{}
				if err := render.Bind(r, data); err != nil {
					return
				}

				todo := *(data.Todo)
				fmt.Println(todo)
			})
		})
	})

	return r
}

type TodoRequest struct {
	*todo.Todo
}

func (t *TodoRequest) Bind(r *http.Request) error {
	if t.Todo == nil {
		return errors.New("missing required Todo fields")
	}
	return nil
}
