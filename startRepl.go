package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// reads input and calls command that user typed in
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
