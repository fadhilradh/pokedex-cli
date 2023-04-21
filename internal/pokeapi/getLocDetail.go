package pokeapi

import (
	"encoding/json"
	"fmt"

	// "fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocDetail(locName string) (LocDetail, error) {
	reqURL := baseURL + fmt.Sprintf("/location-area/%s", locName)

	if val, ok := c.cache.Get(reqURL); ok {
		locDetail := LocDetail{}
		err := json.Unmarshal(val, &locDetail)
		if err != nil {
			return LocDetail{}, err
		}
		fmt.Println("data available : using cache...")
		return locDetail, nil
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return LocDetail{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocDetail{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return LocDetail{}, err
	}

	locDetail := LocDetail{}
	err = json.Unmarshal(body, &locDetail)
	if err != nil {
		return LocDetail{}, err
	}

	c.cache.Add(reqURL, body)
	return locDetail, nil
}
