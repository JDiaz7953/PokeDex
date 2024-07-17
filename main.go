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

//Struct to parse JSON
type config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
		"map" : {
			callback: commandMap,
		},
	}
	
}

func main(){
	commands := DisplayCL()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text, ok := commands[scanner.Text()];
		if !ok {
			fmt.Println("Invalid Command")
			fmt.Println("Try " + " 'help' " + " menu")
			continue
		}
		text.callback()
}
}