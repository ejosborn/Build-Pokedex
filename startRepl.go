package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// reads input and calls command that user typed in
func startRepl(cfg *config) {
	readInput := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		readInput.Scan()

		input := cleanInput(readInput.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		if len(input) == 2 {
			fmt.Println(input[1])
			cfg.userInput = &input[1]
		}

		commandWord, exist := getCommands()[commandName]
		if exist {
			err := commandWord.command(cfg)
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

// takes input and normalizes it for logic
func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}
