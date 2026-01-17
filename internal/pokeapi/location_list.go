package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	//defer resp.Body.Close()
	defer func() {
		if err := resp.Body.Close(); err != nil {
			// Handle or log the error
			log.Printf("Error closing response body: %v", err)
		}
	}()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}

func (c *Client) ListLocationArea(areaName string) (LocationArea, error) {
	if areaName == "" {
		log.Printf("Unspecified area name to retrieve details")
		return LocationArea{}, errors.New("no area specified")
	}
	url := baseURL + "/location-area/" + areaName

	if val, ok := c.cache.Get(url); ok {
		log.Printf("cache hit [%s]", url)
		locationsResp := LocationArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	log.Printf("GET %s returned", url)
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationsResp := LocationArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}
