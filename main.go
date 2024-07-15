package main

import (
	"os"
	"bufio"
	"fmt"
	
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}


// Map that holds the commands the program runs on
func DisplayCL() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	
}

func main(){
	commands := DisplayCL()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text, ok := commands[scanner.Text()]
		if !ok {
			fmt.Println("Invalid Command")
			fmt.Println("Try" + "'help'" + "menu")
		}
		text.callback()




		
	
}
}