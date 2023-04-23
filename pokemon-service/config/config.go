package config

import (
	"encoding/json"
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

func createMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/pokemons", func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := json.Marshal(handlers.GetPokemons())
		w.Write(bytes)
	})

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

	return &Config{Mux: createMux(), Host: host, Port: port}
}
