package main

import (
	"os"
	"bufio"
	"fmt"
	 "github.com/JDiaz7953/PokeDex/internal/pokeapi"
	 "strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.ChangePage, *pokeapi.PokemonInArea,  string) error
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
		"explore" : {
			name: "explore",
			description: "explore area to find pokemon by explore <location>",
			callback: commandExplore,
		},
	}
	
}

func main(){
	commands := DisplayCL()
	scanner := bufio.NewScanner(os.Stdin)


	ChangeP := pokeapi.ChangePage{
		APIClient: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		NextLocationUrl: nil,
	    PrevLocationUrl: nil,
	}

	explore := pokeapi.PokemonInArea{
		LocationAreaURL: "https://pokeapi.co/api/v2/location-area/" ,

	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)
		cmd := parts[0]
		nameOfArea := ""
		if cmd == "explore" && len(parts) > 1 {
			nameOfArea = parts[1]
		}
		text, ok := commands[cmd];
		if !ok {
			fmt.Println("Invalid Command")
			fmt.Println("Try " + " 'help' " + " menu")
			continue
		}
		text.callback(&ChangeP, &explore, nameOfArea  )
}
}