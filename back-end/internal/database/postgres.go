package database

import (
	"context"
	"fmt"
	"github.com/PossibleNPC/sample-pokemon-app/internal/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

// Okay, so the config for these utils needs to be propagated down to here from Main since these utils aren't made to
// be standalone, unlike the ingest_data script.

func ConnPool(context context.Context, cfg *config.Config)(*pgxpool.Pool, error) {
	psqlConnString := fmt.Sprintf("postgres://%s:%s@%s:%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)
	dbpool, err := pgxpool.Connect(context, psqlConnString)
	if err != nil {
		return nil, err
	}
	return dbpool, nil
}

func GetPokemon(context context.Context, dbpool *pgxpool.Pool) (pgx.Rows, error) {
	stmt := `SELECT * from public.pokemon limit 10;`

	rows, err := dbpool.Query(context, stmt)
	if err != nil {
		log.Fatalf("received error: %s\n", err)
	}
	return rows, nil
}