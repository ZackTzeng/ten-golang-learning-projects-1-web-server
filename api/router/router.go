package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
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
			r.Get("/", getAllTodos)
		})
	})

	return r
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all todos"))
}
