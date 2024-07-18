package main
import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"fmt"
	
)

func commandMapBack(ChangeP *ChangePage) error{
	//http Request that gets the Locations. If user called map already it goes to the previous set of 20 questions and so on.
	myUrl := *ChangeP.PrevLocationUrl
	if ChangeP.PrevLocationUrl == nil{
		fmt.Print("PANIC")
		os.Exit(1)
	} // else write error here

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
	var Location LocationAreas
	err = json.Unmarshal(info, &Location)
	if err != nil{
		os.Exit(1)
	}

	//Print Location
	for  _, v := range Location.Results{
		fmt.Println(v.Name)
	}

	ChangeP.NextLocationUrl = Location.Next
	ChangeP.PrevLocationUrl = Location.Next

	return nil
}




// func commandMapBack(ChangeP *ChangePage) error{
// 	Location := GetHttp(ChangeP)

// 	//Print Location
// 	for  _, v := range Location.Results{
// 		fmt.Println(v.Name)
// 	}

// 	ChangeP.NextLocationUrl = Location.Next
// 	ChangeP.PrevLocationUrl = Location.Next

// 	return nil

// 	} 