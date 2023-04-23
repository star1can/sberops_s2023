package service

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"sync"
)

type pokemonServiceImpl struct{}

type pokemonInfo struct {
	Name string `json:"name"`
	Url  string ` json:"url"`
}

type result struct {
	Data  Pokemon
	Error error
}

func (*pokemonServiceImpl) GetPokemons() ([]Pokemon, error) {
	pokemonsInfo, err := getPokemonsInfo()
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	pokemonChnl := make(chan result, len(pokemonsInfo))

	for _, info := range pokemonsInfo {
		wg.Add(1)
		go func(pokemonInfo pokemonInfo) {
			defer wg.Done()
			pokemonChnl <- getPokemonInfo(pokemonInfo)
		}(info)
	}

	wg.Wait()
	close(pokemonChnl)

	pokemonList := make([]Pokemon, 0, len(pokemonChnl))
	for res := range pokemonChnl {
		if res.Error != nil {
			return nil, err
		}
		pokemonList = append(pokemonList, res.Data)
	}

	sortById(pokemonList)

	return pokemonList, nil
}

func performGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func getPokemonsInfo() ([]pokemonInfo, error) {
	body, err := performGetRequest(URL)

	if err != nil {
		return nil, err
	}

	data := struct {
		PokemonsInfo []pokemonInfo `json:"results"`
		Count        int           `json:"count"`
	}{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data.PokemonsInfo, nil
}

func getPokemonInfo(info pokemonInfo) result {
	body, err := performGetRequest(info.Url)
	if err != nil {
		return result{Error: err}
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return result{Error: err}
	}

	return result{Data: pokemon}
}

func sortById(pokemonList []Pokemon) {
	sort.Slice(pokemonList, func(i, j int) bool {
		return pokemonList[i].Id < pokemonList[j].Id
	})
}
