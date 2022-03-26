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

	router := api.Api()
	err := http.ListenAndServe(":" + cfg.Api.Port, router); if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}