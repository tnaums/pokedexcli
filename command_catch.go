package main

import (
	"fmt"
	"math/rand"
	"pokedexcli/internal/pokeapi"
	"pokedexcli/internal/dex"
)

func commandCatch(cfg *config, pokemon string) error {
	pokemonData, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	// Base experience numbers range from 1 to 255
	chance := rand.Intn(255) + 1
	fmt.Printf("Base: %d\n", pokemonData.BaseExperience)
	fmt.Printf("Random: %d\n", chance)
	if pokemonData.BaseExperience > chance {
		fmt.Println("Missed!")
	} else {
		fmt.Println("Caught!")
		newPokemon := shrinkStruct(pokemonData)
		cfg.pokedex[pokemon] = newPokemon
		fmt.Printf("You have caught %d pokemon.\n", len(cfg.pokedex))
	}
	return nil
}

func shrinkStruct(pokemonData pokeapi.Catch) dex.Pokemon {
	fmt.Println(pokemonData.BaseExperience)
	statsStruct := dex.Stats{pokemonData.Stats[0].BaseStat, pokemonData.Stats[1].BaseStat, pokemonData.Stats[2].BaseStat, pokemonData.Stats[3].BaseStat, pokemonData.Stats[4].BaseStat, pokemonData.Stats[5].BaseStat}
	embodimentsList := make([]string, 0)
	for _, x := range pokemonData.Types{
		embodimentsList = append(embodimentsList, x.Type.Name)
	}
	pokemonStruct := dex.Pokemon{pokemonData.Name, pokemonData.Height, pokemonData.Weight, statsStruct, embodimentsList}
	fmt.Println(pokemonStruct)

	return pokemonStruct
}
