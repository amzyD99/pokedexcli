package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("usage: explore <area-name>")
		return nil
	}
	name := args[0]
	fmt.Printf("Exploring %s...\n", name)

	result, err := cfg.pokeapiClient.GetLocationArea(name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, p := range result.PokemonEncounters {
		fmt.Println(" -", p.Pokemon.Name)
	}
	return nil
}
