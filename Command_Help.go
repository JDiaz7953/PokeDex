package main

import (
	"fmt"
)

func commandHelp(ChangeP *ChangePage) error {
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
	fmt.Println("")
	return nil
}
