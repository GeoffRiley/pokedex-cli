package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseURL + endpoint

	dat, ok := c.cache.Get(fullUrl)
	if ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil

	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		fmt.Print("error na new request")
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Print("Error na DO")
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullUrl, data)
	return pokemon, nil
}
