package main

import ("fmt")

func commandHelp() error {
	commands := DisplayCL()
	fmt.Println("Welcome to the Pokedex!")
		    fmt.Println("") 
			fmt.Println("Usage:")
			fmt.Println("") 
			fmt.Println(commands["help"].name + ": " + commands["help"].description )
			fmt.Println(commands["exit"].name + ": " + commands["exit"].description )
			fmt.Println("") 
		return nil
}