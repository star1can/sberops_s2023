package handlers

import (
	"fmt"
	"gitlab.atp-fivt.org/courses/homework-solutions/pokemon-api-client/service"
	"net/http"
	"strings"
)

func PokemonsMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/get-all", func(writer http.ResponseWriter, request *http.Request) {
		pokemons, err := service.GetPokemonService().GetPokemons()
		if err != nil {
			writer.Write([]byte("Ошибка получения покемонов! " + err.Error()))
			return
		}

		writer.Write([]byte(getInfo(pokemons)))
	})

	return mux
}

func getInfo(pokemons []service.Pokemon) string {
	var sb strings.Builder
	sb.WriteString("Доступные покемоны: \n")
	for _, p := range pokemons {
		sb.WriteString(fmt.Sprintf("  -%s\n", p.Name))
	}
	return sb.String()
}
