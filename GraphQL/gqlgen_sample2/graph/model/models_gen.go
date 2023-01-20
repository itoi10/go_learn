// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Battle struct {
	ID        string  `json:"id"`
	PokemonID string  `json:"pokemon_id"`
	Opponent  string  `json:"opponent"`
	Location  *string `json:"location"`
	Date      *string `json:"date"`
}

type NewBattle struct {
	PokemonID string `json:"pokemon_id"`
	Opponent  string `json:"opponent"`
	Location  string `json:"location"`
}

type NewPokemon struct {
	Name        string  `json:"name"`
	Power       *string `json:"Power"`
	Description string  `json:"Description"`
	Date        string  `json:"date"`
}

type Pokemon struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Power       *string `json:"power"`
	Description *string `json:"description"`
}