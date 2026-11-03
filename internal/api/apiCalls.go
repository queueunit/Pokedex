package apiCalls

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	pokecache "github.com/queueunit/Pokedex/internal/cache"
)

const locationAreaURL = "https://pokeapi.co/api/v2/location-area/"

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

func GetDetailedLocation(loc string, cache *pokecache.Cache) (DetailedLocation, error) {

	area := locationAreaURL + loc

	cacheResponse, inCache := cache.Get(area)
	if inCache {
		//		fmt.Println("Found in cache") // For troubleshooting purposes
		results := DetailedLocation{}
		err := json.Unmarshal(cacheResponse, &results)
		if err != nil {
			return DetailedLocation{}, errors.New("Could not unmarshal cached data")
		}
	}

	baseURL, err := url.Parse(area)
	if err != nil {
		panic(err)
	}

	resp, err := http.Get(baseURL.String())
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var location DetailedLocation
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&location)
	if err != nil {
		return DetailedLocation{}, errors.New("Invalid")
	}

	cachedResult, err := json.Marshal(location)
	if err != nil {
		return location, errors.New("can not cache results")
	}

	cache.Add(area, cachedResult)

	return location, errors.New("")
}
