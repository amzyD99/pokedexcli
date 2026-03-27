package main

import "fmt"

func commandMap(cfg *config) error {
	
	result, err := cfg.pokeapiClient.GetLocationAreas(cfg.Next)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cfg.Next = result.Next
	cfg.Previous = result.Previous
	for _, area := range result.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	result, err := cfg.pokeapiClient.GetLocationAreas(cfg.Previous)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cfg.Next = result.Next
	cfg.Previous = result.Previous
	for _, area := range result.Results {
		fmt.Println(area.Name)
	}
	return nil
}
