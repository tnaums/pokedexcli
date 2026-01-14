package main

import "fmt"

func commandInspect(cfg *config, pokemon string) error {
	p, ok := cfg.pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	fmt.Printf("  -hp: %d\n", p.Stats.Hp)
	fmt.Printf("  -attack: %d\n", p.Stats.Attack)
	fmt.Printf("  -defense: %d\n", p.Stats.Defense)
	fmt.Printf("  -special-attack: %d\n", p.Stats.SpecialA)
	fmt.Printf("  -special-defense: %d\n", p.Stats.SpecialD)
	fmt.Printf("  -speed: %d\n", p.Stats.Speed)
	fmt.Println("Types:")

	for _, embodiment := range p.Embodiments {
		fmt.Printf("  - %s\n", embodiment)
	}
	return nil
}
