package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

const locationAreaURL = "https://pokeapi.co/api/v2/location-area/"
const pokemonSearchURL = "https://pokeapi.co/api/v2/pokemon/"

type cliCommand struct {
	name        string
	description string
	command     func(cfg *config) error
}

// returns commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			command:     commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next page of locations",
			command:     commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous page of locations",
			command:     commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location to see what Pokemon there are",
			command:     commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon!",
			command:     commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Program",
			command:     commandExit,
		},
	}
}

// outputs help menu
func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}

// exits program
func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}

// displays next page of locations
func commandMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	fmt.Println()

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

// displays previous page of locations
func commandMapB(cfg *config) error {

	if cfg.prevLocationAreaURL == nil {
		return errors.New("You are at the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	fmt.Println()

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

// lists pokemon found at a location
func commandExplore(cfg *config) error {

	if cfg.userInput == nil {
		return errors.New("Location doesn't exist")
	}

	search := locationAreaURL + *cfg.userInput

	cfg.searchLink = &search

	resp, err := cfg.pokeapiClient.ListLocationAreasPokemon(cfg.searchLink)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s", *cfg.userInput)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	fmt.Println()

	return nil
}

// tries to catch pokemon
func commandCatch(cfg *config) error {
	if cfg.userInput == nil {
		return errors.New("Please enter a pokemon to catch")
	}

	search := pokemonSearchURL + *cfg.userInput

	cfg.searchLink = &search

	resp, err := cfg.pokeapiClient.ListPokemonInfo(cfg.searchLink)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", *cfg.userInput)
	catch := rand.Intn(resp.BaseExperience)

	if catch > 40 {
		fmt.Printf("%s escaped\n", *cfg.userInput)
		return nil
	}

	fmt.Printf("%s was caught\n", *cfg.userInput)

	//add to pokedex map

	return nil
}
