package main

import (
	//	"errors"
	"fmt"
)

func commandExplore(cfg *config, area string) error {
	fmt.Printf("You are exploring area %s!\n", area)
	cfg.pokeapiClient.ListPokemon(area)
	return nil
}
