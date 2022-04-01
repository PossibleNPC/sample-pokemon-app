package main

// TODO: Purpose of this script is to start our server for the API

import (
	"fmt"
	"github.com/PossibleNPC/sample-pokemon-app/internal/config"
	"github.com/PossibleNPC/sample-pokemon-app/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
)

type Application struct {
	cfg *config.Config
	db *database.PokemonDb
	router chi.Router
}

func main() {
	app := &Application{}
	app.cfg = config.NewConfig()
	app.cfg.GetConfFile("./back-end/config.yml")
	app.cfg.GetConfEnv()
	app.cfg.PsqlConnString()
	app.db = database.NewPokemonDb(app.cfg.Database.Dsn)

	app.router = chi.NewRouter()
	app.router.Use(middleware.Logger)
	app.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go-Chi!"))
	})
	app.router.Get("/pokemon", app.HandlePokemonResults)
	app.router.Get("/pokemon/{name}", app.HandleNameResults)
	app.router.Get("/pokemon/{type}", app.HandleTypeResults)

	// TODO: The context for the app needs to come from Main, and not from the sub-modules
	// More research is needed to fully understand Context.
	err := http.ListenAndServe(":" + app.cfg.Api.Port, app.router); if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}