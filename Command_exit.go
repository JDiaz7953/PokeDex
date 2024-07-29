package main

import ("os"
		"fmt"
		"github.com/JDiaz7953/PokeDex/internal/pokeapi"
 )

func commandExit(ChangeP *pokeapi.ChangePage, ex *pokeapi.PokemonInArea,  nameLoc string) error {
	//Exits Program
	fmt.Println("You are now leaving the Pokedex")
	os.Exit(0)
	return nil
}