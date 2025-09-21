package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaResponse, error) {
	endpoint := "/location-area" // corrected from /location_area
	fullUrl := baseUrl + endpoint
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	data, ok := c.cache.Get(fullUrl)
	if ok {
		// cache hit
		fmt.Println("cache hit")
		locationAreaResp := LocationAreaResponse{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationAreaResp, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	locationAreaResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(fullUrl, data)
	return locationAreaResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseUrl + endpoint

	data, ok := c.cache.Get(fullUrl)
	if ok {
		// cache hit
		fmt.Println("cache hit")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	fmt.Println("cache miss")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullUrl, data)
	return locationArea, nil
}
