package main

import (
	"fmt"
	"github.com/JDiaz7953/PokeDex/internal/pokeapi"
)

func commandHelp(ChangeP *pokeapi.ChangePage, ex *pokeapi.PokemonInArea,  nameLoc string) error {
	//Displays Commands
	commands := DisplayCL()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println(commands["help"].name + ": " + commands["help"].description)
	fmt.Println(commands["exit"].name + ": " + commands["exit"].description)
	fmt.Println(commands["map"].name + " : " + commands["map"].description)
	fmt.Println(commands["mapb"].name + ": " + commands["mapb"].description)
	fmt.Println(commands["explore"].name + ": " + commands["explore"].description)
	fmt.Println("")
	return nil
}
