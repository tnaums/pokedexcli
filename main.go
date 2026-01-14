package main

import (
	"time"

	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/dex"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	pokeDex := dex.NewDex()
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: pokeDex,
	}

	startRepl(cfg)
}
