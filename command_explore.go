package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	//input is slice of strings
	if len(args) == 0 {
		return errors.New("no location area provided")
	}
	if len(args) > 1 {
		return errors.New("to many location areas provided, only explore one at once")
	}
	locationAreaName := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocationAreas(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
