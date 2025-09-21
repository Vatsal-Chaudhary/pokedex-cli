package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonname string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonname
	fullUrl := baseUrl + endpoint

	data, ok := c.cache.Get(fullUrl)
	if ok {
		// cache hit
		fmt.Println("cache hit")
		pokemonName := Pokemon{}
		err := json.Unmarshal(data, &pokemonName)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonName, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemonName := Pokemon{}
	err = json.Unmarshal(data, &pokemonName)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)
	return pokemonName, nil
}
