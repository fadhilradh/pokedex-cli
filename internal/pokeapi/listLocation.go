package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Location, error) {
	reqURL := baseURL + "/location-area"
	if pageURL != nil {
		reqURL = *pageURL
	}

	if val, ok := c.cache.Get(reqURL); ok {
		locations := Location{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return Location{}, err
		}
		fmt.Println("data available : using cache...")
		return locations, nil
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return Location{}, err
	}

	locations := Location{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(reqURL, body)
	return locations, nil
}
