package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"sync"
)

const (
	URL string = "https://pokeapi.co/api/v2/pokemon/"
)

type Pokemon struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	BaseExperience int64  `json:"base_experience"`
	Height         int64  `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int64  `json:"order"`
	Weight         int64  `json:"weight"`
}

type pokemonInfo struct {
	Name string `json:"name"`
	Url  string ` json:"url"`
}

type result struct {
	Data  Pokemon
	Error error
}

func performGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func getPokemonsInfo() []pokemonInfo {
	body, err := performGetRequest(URL)

	if err != nil {
		panic(err)
	}

	data := struct {
		PokemonsInfo []pokemonInfo `json:"results"`
		Count        int           `json:"count"`
	}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	return data.PokemonsInfo
}

func GetPokemons() []Pokemon {
	pokemonsInfo := getPokemonsInfo()

	wg := sync.WaitGroup{}
	pokemonChnl := make(chan result, len(pokemonsInfo))
	for _, info := range pokemonsInfo {
		wg.Add(1)
		go func(pokemonInfo pokemonInfo) {
			defer wg.Done()

			body, err := performGetRequest(pokemonInfo.Url)
			if err != nil {
				pokemonChnl <- result{Error: err}
				return
			}

			pokemon := Pokemon{}
			err = json.Unmarshal(body, &pokemon)
			if err != nil {
				pokemonChnl <- result{Error: err}
				return
			}

			pokemonChnl <- result{Data: pokemon, Error: nil}
		}(info)
	}

	wg.Wait()
	close(pokemonChnl)

	pokemonList := make([]Pokemon, 0, len(pokemonChnl))
	for res := range pokemonChnl {
		if res.Error != nil {
			panic(res.Error)
		}
		pokemonList = append(pokemonList, res.Data)
	}

	sort.Slice(pokemonList, func(i, j int) bool {
		return pokemonList[i].Id < pokemonList[j].Id
	})

	return pokemonList
}
