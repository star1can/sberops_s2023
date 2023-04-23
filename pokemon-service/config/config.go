package config

import (
	"gitlab.atp-fivt.org/courses/homework-solutions/pokemon-api-client/handlers"
	"net/http"
	"os"
	"strconv"
)

type Config struct {
	Mux  *http.ServeMux
	Host string
	Port int
}

func getMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/pokemons/", http.StripPrefix("/pokemons", handlers.PokemonsMux()))
	return mux
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func GetConfig() *Config {
	host := getEnv("HOST", "")

	portString := getEnv("PORT", "8083")
	port, err := strconv.Atoi(portString)
	if err != nil {
		panic(err)
	}

	return &Config{Mux: getMux(), Host: host, Port: port}
}
