package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"pokedexcli/internal/dex"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokemapCache pokecache.Cache
	pokedex map[string]dex.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		area := "eterna-city-area"  // default
		if len(words) > 1 {
			area = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, area)
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

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {

			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "list of all the Pokémon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name: "catch",
			description: "attempt to catch",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "inspect a pokemon from the pokedex",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "list all captured pokemon in the pokedex",
			callback: commandPokedex,
		},
	}
}
