package main

import (
	"fmt"
	"log"
	"os"
)

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
		"exit": {
			name:        "exit",
			description: "Exits the Program",
			command:     commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays some locations",
			command:     commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous locations",
			command:     commandMapB,
		},
	}
}

// outputs
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

// displays next locations
func commandMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Print("- %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

// displays previous locations
func commandMapB(cfg *config) error {

	return nil
}
