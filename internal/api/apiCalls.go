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

func GetLocations(offset int) (PokeLocations, error) {

	// Hardcoded url for now, maybe pass in as variable later?
	baseURL, err := url.Parse("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		panic(err)
	}

	// Add the query parameter "offset" to paginate
	params := url.Values{}
	params.Add("offset", string(offset))

	// Encode and attatch parameters to the URL
	baseURL.RawQuery = params.Encode()

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
