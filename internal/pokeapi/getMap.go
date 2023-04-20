package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Maps, error) {
	reqURL := baseURL + "/location-area"
	if pageURL != nil {
		reqURL = *pageURL
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return Maps{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Maps{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return Maps{}, err
	}

	maps := Maps{}
	err = json.Unmarshal(body, &maps)
	if err != nil {
		return Maps{}, err
	}

	return maps, nil
}
