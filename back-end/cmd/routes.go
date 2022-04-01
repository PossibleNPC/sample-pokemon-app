package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go-Chi!"))
	})
	router.Get("/pokemon", app.HandlePokemonResults)
	router.Get("/pokemon/{name}", app.HandleNameResults)
	router.Get("/pokemon/{type}", app.HandleTypeResults)

	return router
}