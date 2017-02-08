package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gocraft/web"
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
	panicOnError(http.ListenAndServe("0.0.0.0:"+config.Port, router))
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

func jsonParse(req *web.Request, target interface{}) (interface{}, error) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&target)
	defer req.Body.Close()
	return target, err
}

func jsonAnswer(rw web.ResponseWriter, data interface{}) {
	if rw.Written() == false {
		data, err := json.Marshal(data)
		jsonError(rw, err)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(data)
	}
}

func jsonError(rw web.ResponseWriter, err error) {
	if err != nil {
		jData, _ := json.Marshal(jsn{
			"Error": err.Error(),
		})
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(jData)
	}
}

type jsn map[string]string
