package main

type Pokedex struct {
	pokedex map[string]Pokemon
}

type Pokemon struct {
	name   string
	height int
	weight int
	stats  PokemonStats
	typing PokemonTypes
}

type PokemonStats struct {
	hp             int
	attack         int
	defense        int
	specialAttack  int
	specialDefense int
	speed          int
}

type PokemonTypes struct {
	primaryT   string
	secondaryT string
}

func NewPokemon() {

}

func AddPokemon() {

}
