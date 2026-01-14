package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

//ListPokemon
func (c *Client) ListPokemon(location string) (RespLocationArea, error){
	url := "https://pokeapi.co/api/v2/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		listPokemonResp := RespLocationArea{}
		err := json.Unmarshal(val, &listPokemonResp)
		if err != nil {
			return RespLocationArea{}, err
		}

		return listPokemonResp, nil
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	listPokemonResp := RespLocationArea{}
	err = json.Unmarshal(dat, &listPokemonResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, dat)
	return listPokemonResp, nil
	
}
