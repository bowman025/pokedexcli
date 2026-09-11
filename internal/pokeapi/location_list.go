package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type LocationResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *Client) GetLocationList(urlAddress *string) (LocationResponse, error) {
	urlValue := baseUrl + "/location-area"
	if urlAddress != nil {
		urlValue = *urlAddress
	}

	if data, ok := c.cache.Get(urlValue); ok {
		pokeRes := LocationResponse{}
		err := json.Unmarshal(data, &pokeRes)
		if err != nil {
			return LocationResponse{}, fmt.Errorf("error during unmarshal: %v", err)
		}

		return pokeRes, nil
	}

	res, err := c.httpClient.Get(urlValue)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("API error: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error reading the request: %v", err)
	}

	pokeRes := LocationResponse{}
	err = json.Unmarshal(data, &pokeRes)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error during unmarshal: %v", err)
	}

	c.cache.Add(urlValue, data)

	return pokeRes, nil
}
