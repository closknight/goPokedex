package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/closknight/goPokedex/internal/client"
)

type config struct {
	next   *string
	prev   *string
	client client.Client
}

func parse(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	cacheInterval := 5 * time.Minute

	config := config{
		next:   nil,
		prev:   nil,
		client: client.NewClient(cacheInterval),
	}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		tokens := parse(scanner.Text())
		if len(tokens) == 0 {
			continue
		}

		cmdName := tokens[0]
		args := tokens[1:]
		command, exists := commands[cmdName]
		if exists {
			err := command.callback(&config, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Wrong command given. use help for guidance.")
		}
	}
}
