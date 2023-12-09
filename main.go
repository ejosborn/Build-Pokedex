package main

import (
	"time"

	"github.com/ejosborn/Build-Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	userInput           *string
	searchLink          *string
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := config{
		pokeapiClient: pokeClient,
	}

	startRepl(&cfg)
}
