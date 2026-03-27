package main

import "fmt"

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	return fmt.Errorf("exit")
}