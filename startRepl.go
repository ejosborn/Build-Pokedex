package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	readInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		readInput.Scan()

		input := cleanInput(readInput.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		command, exist := getCommands()[commandName]
		if exist {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// returns commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Program",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
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

func commandExit() error {
	os.Exit(0)
	return nil
}
