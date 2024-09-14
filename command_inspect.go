package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	//input is slice of strings
	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	}
	if len(args) > 1 {
		return errors.New("to many names provided, or maybe a space")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		err_msg := fmt.Sprintf("%s not caught yet", pokemonName)
		return errors.New(err_msg)
	}

	//pokemon in dex so lets print some info
	fmt.Printf("\nName: %s\n", pokemon.Name)
	fmt.Printf("Pokedex ID: %v\n", pokemon.ID)
	fmt.Printf("Types: ")
	for _, typ := range pokemon.Types {
		fmt.Printf("%s, ", typ.Type.Name)
	}
	fmt.Printf("\nHeight: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s:%v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("")
	return nil
}
