package main

// TODO: Purpose of this script is to start our server for the API

import (
	"github.com/PossibleNPC/sample-pokemon-app/internal/config"
	"github.com/PossibleNPC/sample-pokemon-app/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

type application struct {
	cfg *config.Config
	errorLog *log.Logger
	infoLog *log.Logger
	db *database.PokemonDb
	router chi.Router
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{errorLog: errorLog, infoLog: infoLog}
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
		errorLog.Fatalf("failed to start the server: %s", err)
	}
}