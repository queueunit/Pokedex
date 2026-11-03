package apiCalls

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	pokecache "github.com/queueunit/Pokedex/internal/cache"
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
func GetLocations(endpoint string, cache *pokecache.Cache) (PokeLocations, error) {

	cacheResponse, inCache := cache.Get(endpoint)
	if inCache {
		//		fmt.Println("Found in cache") // For troubleshooting purposes
		results := PokeLocations{}
		err := json.Unmarshal(cacheResponse, &results)
		if err != nil {
			return PokeLocations{}, errors.New("Could not unmarshal cached data")
		}
	}

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

	cachedResult, err := json.Marshal(locations)
	if err != nil {
		return locations, errors.New("can not cache results")
	}

	cache.Add(endpoint, cachedResult)

	return locations, errors.New("")
}
