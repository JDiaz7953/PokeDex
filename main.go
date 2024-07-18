package main

import (
	"os"
	"bufio"
	"fmt"
	
)

type cliCommand struct {
	name        string
	description string
	callback    func(*ChangePage) error
}

type ChangePage struct {
	APIClient string
	NextLocationUrl *string
	PrevLocationUrl *string
}


//Struct to parse JSON
type LocationAreas struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Prev *string   `json:"previous"`
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
			name:        "map",
			description: "Displays list of locations",
			callback: commandMap,
		},

		"mapb" : {
			name:        "mapb",
			description: "Displays previous list of locations",
			callback: commandMapBack,
		},
	}
	
}

func main(){
	commands := DisplayCL()
	scanner := bufio.NewScanner(os.Stdin)
	ChangeP := ChangePage{
		APIClient: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		NextLocationUrl: nil,
	    PrevLocationUrl: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text, ok := commands[scanner.Text()];
		if !ok {
			fmt.Println("Invalid Command")
			fmt.Println("Try " + " 'help' " + " menu")
			continue
		}
		text.callback(&ChangeP)
}
}