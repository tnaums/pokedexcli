package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

//Catch Pokemon
func (c *Client) CatchPokemon(name string) (Catch, error){
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pokemonData := Catch{}
		err := json.Unmarshal(val, &pokemonData)
		if err != nil {
			return Catch{}, err
		}

		return pokemonData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Catch{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Catch{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Catch{}, err
	}
	pokemonData := Catch{}
	err = json.Unmarshal(dat, &pokemonData)
	if err != nil {
		return Catch{}, err
	}

	c.cache.Add(url, dat)
	return pokemonData, nil
	
}
