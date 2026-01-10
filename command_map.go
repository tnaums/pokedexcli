package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	type Location struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}
	
	type MapLoc struct {
    Count int `json:"count"` 
    Next *string `json:"next"` 
		Previous *string `json:"previous"`
		Results []Location `json:"results"`
	}
	
	locs := MapLoc{}
	newerr := json.Unmarshal(body, &locs)
	if newerr != nil {
		fmt.Println(newerr)
	}
	//	fmt.Printf("%s", body)
	// after Unmarshal succeeds
for _, loc := range locs.Results {
    fmt.Println(loc.Name)
}
	return nil
}
