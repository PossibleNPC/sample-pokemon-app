package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PossibleNPC/sample-pokemon-app/internal/config"
	"github.com/jackc/pgx/v4"
)


func main() {
	var cfg config.Config

	config.GetConfFile(&cfg, "./back-end/config.yml")
	config.GetConfEnv(&cfg)

	psqlConnString := fmt.Sprintf("postgres://%s:%s@%s:%d",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)

	conn, err := pgx.Connect(context.Background(), psqlConnString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to the database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// TODO: This just opens our CSV. We need to increase the flexibility of this tool. Not sure yet if I want to do
	// it through a CLI, or just have a script run.
	inputFile, err := os.Open("../data/pokemon_data_edited.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	res, err := conn.PgConn().CopyFrom(context.Background(), inputFile, "COPY pokemon FROM STDIN CSV HEADER")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error copying CSV: %v data to database.", inputFile.Name())
		os.Exit(1)
	}
	fmt.Println(res.RowsAffected())
}
