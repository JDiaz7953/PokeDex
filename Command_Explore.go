package main

import(
	"os"
	"github.com/JDiaz7953/PokeDex/internal/pokeapi"
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)



func commandExplore(ChangeP *pokeapi.ChangePage, explore *pokeapi.PokemonInArea, nameLoc string) error {
	myUrl := explore.LocationAreaURL + nameLoc
	
	response, err := http.Get(myUrl)
	if err != nil {
		os.Exit(1)
	}
	defer response.Body.Close()

	//read response from the body
	info, err := io.ReadAll(response.Body)
		if err != nil{
			os.Exit(1)
		}
	
	//Unmarshall data(turn from bytes to strings)
	var area pokeapi.Explorelocation
	err = json.Unmarshal(info, &area )
	if err != nil{
		os.Exit(1)
	}

	fmt.Println("Pokemon in this Area: ")
	for  _, v := range area.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}


	return nil
}