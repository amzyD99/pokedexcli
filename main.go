package main

import (
	"time"
	"github.com/amzyd99/pokedexcli/internal/pokeapi"
)

type config struct {
	Next          string
	Previous      string
	pokeapiClient pokeapi.Client
}



func main(){
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	StartRepl(cfg)
}
