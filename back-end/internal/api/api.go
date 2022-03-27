package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


// Api is a function that returns a Go-Chi Mux that has been preconfigured specifically for this application.
// It includes desired middleware, routes, errors, etc. as desired.
func Api() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go-Chi!"))
	})
	router.Get("/pokemon", HandlePokemonResults)
	router.Get("/pokemon/{name}", HandleNameResults)
	router.Get("/pokemon/{type}", HandleTypeResults)
	return router
}