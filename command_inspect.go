package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("usage: inspect <pokemon-name>")
		return nil
	}
	name := args[0]

	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("ID: %d\n", pokemon.ID)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
	}

	return nil
}
