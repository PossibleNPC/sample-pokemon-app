package models

import (
	"fmt"
	"net/http"
)

// TODO: I need to figure out how to better marshall pokemon from the DB to the struct; currently requires us to marshal
// 47!!! different fields ourselves. There has to be a better way to do this.
type Pokemon struct {
	Id            int    `json:"id,omitempty"`
	PokedexNumber int    `json:"pokedexNumber,omitempty"`
	Name       string `json:"name,omitempty"`
	Generation int    `json:"generation,omitempty"`
	Status  string `json:"status,omitempty"`
	Species string `json:"species,omitempty"`
	Type1   string  `json:"type1,omitempty"`
	Type2    string  `json:"type2,omitempty"`
	HeightM  float32 `json:"heightM,omitempty"`
	WeightKg float32 `json:"weightKg,omitempty"`
	Ability1      string `json:"ability1,omitempty"`
	Ability2      string `json:"ability2,omitempty"`
	AbilityHidden string `json:"abilityHidden,omitempty"`
	TotalPoints int `json:"totalPoints,omitempty"`
	Hp      int `json:"hp,omitempty"`
	Attack   int `json:"attack,omitempty"`
	Defense   int `json:"defense,omitempty"`
	SpAttack  int `json:"spAttack,omitempty"`
	SpDefense int `json:"spDefense,omitempty"`
	Speed          int `json:"speed,omitempty"`
	CatchRate      int `json:"catchRate,omitempty"`
	BaseFriendship int    `json:"baseFriendship,omitempty"`
	BaseExperience int    `json:"baseExperience,omitempty"`
	GrowthRate string `json:"growthRate,omitempty"`
	EggType1       string  `json:"eggType1,omitempty"`
	EggType2         string  `json:"eggType2,omitempty"`
	PercentageMale   float32 `json:"percentageMale,omitempty"`
	PercentageFemale float32 `json:"percentageFemale,omitempty"`
	EggCycles     int     `json:"eggCycles,omitempty"`
	AgainstNormal float32 `json:"againstNormal,omitempty"`
	AgainstFire     float32 `json:"againstFire,omitempty"`
	AgainstWater    float32 `json:"againstWater,omitempty"`
	AgainstElectric float32 `json:"againstElectric,omitempty"`
	AgainstGrass float32 `json:"againstGrass,omitempty"`
	AgainstIce    float32 `json:"againstIce,omitempty"`
	AgainstFight  float32 `json:"againstFight,omitempty"`
	AgainstPoison float32 `json:"againstPoison,omitempty"`
	AgainstGround  float32 `json:"againstGround,omitempty"`
	AgainstFlying  float32 `json:"againstFlying,omitempty"`
	AgainstPsychic float32 `json:"againstPsychic,omitempty"`
	AgainstBug   float32 `json:"againstBug,omitempty"`
	AgainstRock   float32 `json:"againstRock,omitempty"`
	AgainstGhost  float32 `json:"againstGhost,omitempty"`
	AgainstDragon float32 `json:"againstDragon,omitempty"`
	AgainstDark  float32 `json:"againstDark,omitempty"`
	AgainstSteel float32 `json:"againstSteel,omitempty"`
	AgainstFairy float32 `json:"againstFairy,omitempty"`
}

type PokemonList struct {
	Pokemon []Pokemon `json:"pokemon,omitempty"`
}

func (pokemon *Pokemon) Bind(r *http.Request) error {
	if pokemon.Name == "" {
		return fmt.Errorf("pokemon name is a required field")
	}
	return nil
}