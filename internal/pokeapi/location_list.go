package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)



func (c *Client) GetLocationAreas(pageUrl string) (locationArea, error) {
	url := baseURL + "/location-area"
	if pageUrl != "" {
		url = pageUrl
	}

	if cached, ok := c.cache.Get(url); ok {
		var result locationArea
		err := json.Unmarshal(cached, &result)
		if err != nil {
			return locationArea{}, err
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationArea{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationArea{}, err
	}

	c.cache.Add(url, body)

	var result locationArea
	err = json.Unmarshal(body, &result)
	if err != nil {
		return locationArea{}, err
	}

	return result, nil
}
