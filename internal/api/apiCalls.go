package apiCalls

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type PokeLocations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Takes a url from the state.Next/state.Previous container
func GetLocations(endpoint string) (PokeLocations, error) {

	// Hardcoded url for now, maybe pass in as variable later?
	baseURL, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	resp, err := http.Get(baseURL.String())
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var locations PokeLocations
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return PokeLocations{}, err
	}

	return locations, errors.New("")
}
