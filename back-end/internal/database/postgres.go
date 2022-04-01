package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PokemonDb struct {
	conn *sql.DB
}

func NewPokemonDb(dsn string) *PokemonDb {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("error opening database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging database: %s", err)
	}

	return &PokemonDb{conn: db}
}

// Okay, so the config for these utils needs to be propagated down to here from Main since these utils aren't made to
// be standalone, unlike the ingest_data script.

func (pokemonDb *PokemonDb) GetPokemon() (*sql.Rows, error) {
	stmt := `SELECT unique_id, name, status, generation from public.pokemon limit 10;`

	rows, err := pokemonDb.conn.Query(stmt)
	if err != nil {
		log.Fatalf("received error: %s\n", err)
	}

	return rows, nil
}