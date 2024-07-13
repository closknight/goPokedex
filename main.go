package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		text = strings.ToLower(text)
		tokens := strings.Fields(text)
		if len(tokens) == 0 {
			continue
		}

		cmdName := tokens[0]
		command, ok := commands[cmdName]
		if !ok {
			fmt.Println("Wrong command given. use help for guidance.")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
