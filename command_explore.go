package main

import (
	//	"errors"
	"fmt"
)

func commandExplore(cfg *config, area string) error {
	listPokemonResp, err := cfg.pokeapiClient.ListPokemon(area)
	if err != nil {
		return err
	}
	
	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, x := range listPokemonResp.PokemonEncounters{
		fmt.Printf(" - %s\n", x.Pokemon.Name)
	}

	return nil
}
