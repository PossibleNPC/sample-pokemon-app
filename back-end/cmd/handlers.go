package main

import (
	"encoding/json"
	"fmt"
	"github.com/PossibleNPC/sample-pokemon-app/internal/models"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

// Routes List

// GET
// I don't think this is a good API format for this. Need to research good API route creation, but we can change that
// out later too since there's not a lot of routes for now
// localhost:XXXX/pokemon/{id}
// localhost:XXXX/pokemon/{name}
// localhost:XXXX/pokemon/{type}
// All Pokemon
// - Paginated results so as to not render all the data
// By Name
// By Id

// Maybe further down the road, we'll add in user account creation, but not really needed now.

// HandleIndex is meant to return the Index HTML page for the app. I'm not sure if it can serve up a React App or not,
// but we will try.
func (app *application) HandleIndex(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("This is the Index HTML Page... Not yet implemented."))
}

// HandlePokemonResults is meant to return results for all Pokemon in the database as paginated results, starting with
// 50 results.
func (app *application) HandlePokemonResults(w http.ResponseWriter, r *http.Request) {
	rows, err := app.db.GetPokemon()

	if err != nil {
		log.Fatalf("%s", err)
	}

	defer rows.Close()

	for rows.Next() {
		//values, err := rows.Values()
		//if err != nil {
		//	log.Fatal("error on returned rows")
		//}
		//
		//res := fmt.Sprintf("%s %s %s %s %s", values[0], values[2], values[3], values[4], values[5])
		pokemon := &models.Pokemon{}

		err = rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Status, &pokemon.Generation)
		if err != nil {
			log.Fatalf("unable to load results: %s", err)
		}
		out, err := json.Marshal(pokemon)
		if err != nil {
			panic(err)
		}

		w.Write(out)
	}

	//w.Write([]byte("Response with all Pokemon and default paginated results."))
}

// HandleTypeResults returns paginated results matching the type of Pokemon a user wants to query, starting with 50 results.
func (app *application) HandleTypeResults(w http.ResponseWriter, r *http.Request) {
	typeParam := chi.URLParam(r, "type")
	w.Write([]byte(fmt.Sprintf("Response with Pokemon Results matching Type; paginated.\n%s", typeParam)))
}

// HandleNameResults returns a result, or results that match a given name. A name can have multiple results because of
// the regional versions added in later gens.
func (app *application) HandleNameResults(w http.ResponseWriter, r *http.Request) {
	nameParam := chi.URLParam(r, "name")
	w.Write([]byte(fmt.Sprintf("Response with either one result, or many in the case of matching same name, but different regions.\n%s", nameParam)))
}