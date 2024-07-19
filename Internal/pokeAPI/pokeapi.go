package pokeapi

import (

)
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


// func GetHttp(ChangeP *ChangePage) LocationAreas{
// 	//http Request that gets the Locations. If user called map already it goes to the next set of 20 questions and so on.
// 	myUrl := ChangeP.APIClient
// 	if ChangeP.NextLocationUrl != nil{
// 		myUrl = *ChangeP.NextLocationUrl
// 	}
// 	response, err := http.Get(myUrl)
// 	if err != nil {
// 		os.Exit(1)
// 	}
// 	defer response.Body.Close()

// 	//read response from the body
// 	info, err := io.ReadAll(response.Body)
// 		if err != nil{
// 			os.Exit(1)
// 		}
	
// 	//Unmarshall data(turn from bytes to strings)
// 	var locations LocationAreas
// 	err = json.Unmarshal(info, &locations )
// 	if err != nil{
// 		os.Exit(1)
// 	}
// 	return locations 
// }
