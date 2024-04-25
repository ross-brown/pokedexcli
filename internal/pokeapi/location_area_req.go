package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	//check the cache
	dat, ok := c.cache.Get(url)
	if ok {
		fmt.Println("CACHE HIT!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("Cache MISSED")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(url, dat)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationAreaName

	//check the cache
	dat, ok := c.cache.Get(url)
	if ok {
		fmt.Println("CACHE HIT!")
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}
	fmt.Println("Cache MISSED")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	dat, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, dat)

	return locationArea, nil
}
