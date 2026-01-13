package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

//ListPokemon
func (c *Client) ListPokemon(location string) {
	url := "https://pokeapi.co/api/v2/location-area/" + location
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	listPokemonResp := RespLocationArea{}
	err = json.Unmarshal(dat, &listPokemonResp)

	fmt.Println(listPokemonResp.PokemonEncounters[0].Pokemon.Name)
	
}
