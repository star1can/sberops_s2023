package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gitlab.atp-fivt.org/courses/homework-solutions/pokemon-api-client/config"
	"log"
	"net/http"
	"strconv"
)

func serve(conf *config.Config) {
	fmt.Println("Starting app at PORT " + strconv.Itoa(conf.Port))
	err := http.ListenAndServe(conf.Host+":"+strconv.Itoa(conf.Port), conf.Mux)
	defer fmt.Println("Finishing app...")
	if err != nil {
		return
	}
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.GetConfig()
	serve(conf)
}
