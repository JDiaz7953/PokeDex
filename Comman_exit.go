package main

import ("os"
		"fmt"
 )

func commandExit(ChangeP *ChangePage) error {
	//Exits Program
	fmt.Println("You are now leaving the Pokedex")
	os.Exit(0)
	return nil
}