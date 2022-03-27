package main

// TODO: Purpose of this script is to start our server for the API

import (
	"fmt"
	"github.com/PossibleNPC/sample-pokemon-app/internal/api"
	"github.com/PossibleNPC/sample-pokemon-app/internal/config"
	"net/http"
	"os"
)


func main() {
	var cfg config.Config
	config.GetConfFile(&cfg, "./back-end/config.yml")
	config.GetConfEnv(&cfg)

	// TODO: The context for the app needs to come from Main, and not from the sub-modules
	// More research is needed to fully understand Context.
	router := api.Api()
	err := http.ListenAndServe(":" + cfg.Api.Port, router); if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}