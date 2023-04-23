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
	log.Fatal(http.ListenAndServe(conf.Host+":"+strconv.Itoa(conf.Port), conf.Mux))
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.GetConfig()
	serve(conf)
}
