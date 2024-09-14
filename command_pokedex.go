package main

import (
	"fmt"
	"sort"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("No Pokémon in the Pokedex yet.")
		return nil
	}

	// Extract keys (IDs) from the map and convert them to integers for sorting
	var ids []int
	for _, pokemon := range cfg.caughtPokemon {
		ids = append(ids, pokemon.ID)
	}

	// Sort the slice of IDs
	sort.Ints(ids)

	// Print the Pokémon sorted by their IDs
	fmt.Printf("\nPokémon in Pokedex:\n")
	for _, id := range ids {
		// Find the corresponding Pokémon by ID
		for _, pokemon := range cfg.caughtPokemon {
			if pokemon.ID == id {
				fmt.Printf("#%v - %s\n", pokemon.ID, pokemon.Name)
			}
		}
	}
	fmt.Println("")

	return nil
}
