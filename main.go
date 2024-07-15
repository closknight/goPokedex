package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	config := config{next: nil, prev: nil}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		tokens := parse(scanner.Text())
		if len(tokens) == 0 {
			continue
		}

		cmdName := tokens[0]
		command, exists := commands[cmdName]
		if exists {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Wrong command given. use help for guidance.")
		}
	}
}
