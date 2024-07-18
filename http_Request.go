package main
import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	)


func GetHttp(ChangeP *ChangePage) LocationAreas{
	//http Request that gets the Locations. If user called map already it goes to the next set of 20 questions and so on.
	myUrl := ChangeP.APIClient
	if ChangeP.NextLocationUrl != nil{
		myUrl = *ChangeP.NextLocationUrl
	}
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
	var Configuration LocationAreas
	err = json.Unmarshal(info, &Configuration)
	if err != nil{
		os.Exit(1)
	}
	return Configuration
}