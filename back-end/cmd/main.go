package main

// TODO: Purpose of this script is to start our server for the API

import (
	"fmt"
	"github.com/PossibleNPC/sample-pokemon-app/internal/config"
)

func main() {
	var cfg config.Config
	config.GetConfFile(&cfg)
	config.GetConfEnv(&cfg)
	fmt.Printf("%+v", cfg)
}