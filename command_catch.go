package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	//input is slice of strings
	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	}
	if len(args) > 1 {
		return errors.New("to many names provided, or maybe a space")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	// standin for catch difficulty
	randNum := rand.Intn(pokemon.BaseExperience)
	const threshold = 60
	fmt.Printf("odds: %d in %d, your roll: %d \n", threshold, pokemon.BaseExperience, randNum)
	if randNum > threshold {
		return fmt.Errorf("%s was not caught, how tragic", pokemonName)
	}
	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
