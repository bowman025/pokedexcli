package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

type PokeResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *Client) GetPokeResponse(urlAddress *string) (PokeResponse, error) {
	urlValue := "https://pokeapi.co/api/v2/location-area"
	if urlAddress != nil {
		urlValue = *urlAddress
	}

	if data, ok := c.cache.Get(urlValue); ok {
		pokeRes := PokeResponse{}
		err := json.Unmarshal(data, &pokeRes)
		if err != nil {
			return PokeResponse{}, fmt.Errorf("error during unmarshal: %v", err)
		}

		return pokeRes, nil
	}

	res, err := c.httpClient.Get(urlValue)
	if err != nil {
		return PokeResponse{}, fmt.Errorf("API error: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeResponse{}, fmt.Errorf("error reading the request: %v", err)
	}

	pokeRes := PokeResponse{}
	err = json.Unmarshal(data, &pokeRes)
	if err != nil {
		return PokeResponse{}, fmt.Errorf("error during unmarshal: %v", err)
	}

	c.cache.Add(urlValue, data)

	return pokeRes, nil
}
