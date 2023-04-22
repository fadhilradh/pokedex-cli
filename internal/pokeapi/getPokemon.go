package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	reqURL := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(reqURL); ok {
		locDetail := Pokemon{}
		err := json.Unmarshal(val, &locDetail)
		if err != nil {
			return Pokemon{}, err
		}
		return locDetail, nil
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return Pokemon{}, err
	}

	locDetail := Pokemon{}
	err = json.Unmarshal(body, &locDetail)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(reqURL, body)
	return locDetail, nil
}
