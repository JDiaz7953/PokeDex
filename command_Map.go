package main

import (
	"fmt"
)

//Prints a page of 20 Locations
func commandMap(ChangeP *ChangePage) error{
	Location := GetHttp(ChangeP)
	for  _, v := range Location.Results{
		fmt.Println(v.Name)
	}
	//Allows us to go to the next or previous set of locations
	ChangeP.NextLocationUrl = Location.Next
	ChangeP.PrevLocationUrl = Location.Prev
	return nil

	} 
	
