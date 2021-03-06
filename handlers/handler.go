package handlers

import (
    "net/http"
    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "github.com/aicyp/escape-dan-back/controllers"
)

var dbInstance controllers.Database

func NewHandler(db controllers.Database) http.Handler {
    router := chi.NewRouter()
    dbInstance = db
    router.MethodNotAllowed(methodNotAllowedHandler)
    router.NotFound(notFoundHandler)
    router.Route("/users", users)
    return router
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