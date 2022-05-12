package api

import (
	"net/http"
	"todo/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var dbInstance database.Database

func NewHandler(db database.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/employee", Employee)
	return router
}

func Employee(router chi.Router) {
	router.Get("/", getAllEmployees)
	router.Post("/", createEmployee)
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/detail", getEmployee)
		router.Put("/update", updateEmployee)
		router.Delete("/delete", deleteEmployee)
	})
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
