package main

import (
	"fmt"
	"github.com/JDiaz7953/PokeDex/internal/pokeapi"
	
)

//Prints a page of 20 Locations
func commandMap(ChangeP *pokeapi.ChangePage, ex *pokeapi.PokemonInArea,  nameLoc string) error{
	Location := pokeapi.GetHttpReq(ChangeP)
	for  _, v := range Location.Results{
		fmt.Println(v.Name)
	}

	//Allows us to go to the next or previous set of locations
	ChangeP.NextLocationUrl = Location.Next
	ChangeP.PrevLocationUrl = Location.Prev
	return nil

	} 
	
