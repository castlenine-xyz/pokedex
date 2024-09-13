package main

import "github.com/castlenine-xyz/pokedex/internal/pokeapi"

// "fmt"

// holds state info
type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)
}
