package main

// TODO: Purpose of this script is to start our server for the API

import (
	"flag"
	"github.com/PossibleNPC/sample-pokemon-app/internal/config"
	"github.com/PossibleNPC/sample-pokemon-app/internal/database"
	"github.com/go-chi/chi/v5"
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
	yamlFile := flag.String("yamlFile", "./back-end/config.yml", "file path to yaml file")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{errorLog: errorLog, infoLog: infoLog}
	app.cfg = config.NewConfig(*yamlFile)
	app.db = database.NewPokemonDb(app.cfg.Database.Dsn)

	app.router = app.routes()

	err := http.ListenAndServe(":" + app.cfg.Api.Port, app.router); if err != nil {
		errorLog.Fatalf("failed to start the server: %s", err)
	}
}