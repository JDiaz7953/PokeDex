package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"encoding/json"
)

//Maybe Http function to pass in next and prev?


func commandMap() error{
	//http Request
	myUrl := "https://pokeapi.co/api/v2/location-area/"
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
	var Configuration config
	err = json.Unmarshal(info, &Configuration)
	if err != nil{
		os.Exit(1)
	}

	//Print Location
	for i := 0; i <20; i++{
		fmt.Println(Configuration.Results[i].Name)

	}

	return nil

	} 
	
