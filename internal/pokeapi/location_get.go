package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(name string) (locationAreaDetail, error) {
	url := baseURL + "/location-area/" + name

	if cached, ok := c.cache.Get(url); ok {
		var result locationAreaDetail
		err := json.Unmarshal(cached, &result)
		if err != nil {
			return locationAreaDetail{}, err
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaDetail{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaDetail{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreaDetail{}, err
	}

	c.cache.Add(url, body)

	var result locationAreaDetail
	err = json.Unmarshal(body, &result)
	if err != nil {
		return locationAreaDetail{}, err
	}

	return result, nil
}
