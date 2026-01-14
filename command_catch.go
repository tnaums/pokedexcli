package main

import "fmt"

func commandCatch(cfg *config, pokemon string) error {
	_, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Trying to catch %s\n", pokemon)
	//	fmt.Println(pokemonData)
	return nil
}
