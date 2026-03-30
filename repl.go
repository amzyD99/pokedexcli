package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func StartRepl(cfg *config){
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		cleanText := cleanInput(scanner.Text())
		inputCommand := cleanText[0]

		cmd, ok := commands[inputCommand]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			if err := cmd.callback(cfg, cleanText[1:]); err != nil {
				os.Exit(0)
			}
		}
	}
}

func cleanInput(text string) []string{
	//string split into words and lowercased
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Display a help message",
            callback:    commandHelp,
        },
		"map": {
			name:        "map",
			description: "Display 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display 20 previous location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "List all Pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
    }
}
