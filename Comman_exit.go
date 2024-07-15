package main

import ("os"
		"fmt"
 )

func commandExit() error {

	fmt.Println("You are now leaving the Pokedex")
	os.Exit(0)
	return nil
}