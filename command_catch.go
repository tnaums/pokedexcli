package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemon string) error {
	pokemonData, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Trying to catch %s\n", pokemon)
	fmt.Printf("Base experience %d\n", pokemonData.BaseExperience)
	// 1 to 255
	chance := rand.Intn(254) + 1
	fmt.Printf("Base: %d\n", pokemonData.BaseExperience)
	fmt.Printf("Random: %d\n", chance)
	if pokemonData.BaseExperience > chance {
		fmt.Println("Missed!")
	} else {
		fmt.Println("Caught!")
	}
	return nil
}
