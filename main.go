package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	Port         string `json:"port"`
	StaticFolder string `json:"staticFolder"`
}

var config Config

type GlobalContext struct{}

func main() {
	loadConfig()

	router := setRouter()

	fmt.Println("Server listening at ", config.Port)
	panicOnError(http.ListenAndServe("localhost:"+config.Port, router))
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func loadConfig() {
	configFile, err := os.Open("config.json")
	if err != nil {
		panicOnError(err)
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		panicOnError(err)
	}
}
