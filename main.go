package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"

	"github.com/gocraft/web"
)

var config Config

func main() {
	loadConfig()
	router := setRouter()

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("galax.be"),
		Cache:      autocert.DirCache("certs"),
	}

	fmt.Println("Server listening at ", config.Port)
	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	server.ListenAndServeTLS("", "")
	//panicOnError(http.ListenAndServe("0.0.0.0:"+config.Port, router))
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO: Handle config file inexistence
func loadConfig() {
	configFile, err := os.Open("config.json")
	panicOnError(err)

	jsonParser := json.NewDecoder(configFile)
	err1 := jsonParser.Decode(&config)
	panicOnError(err1)
}

func jsonReqParse(req *web.Request, target interface{}) (interface{}, error) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&target)
	defer req.Body.Close()
	return target, err
}

func jsonParse(data string) jsn {
	s := jsn{}
	err := json.Unmarshal([]byte(data), &s)
	panicOnError(err)
	return s
}

func jsonAnswer(rw web.ResponseWriter, data interface{}) {
	if rw.Written() == false {
		data, err := json.Marshal(data)
		jsonError(rw, err)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(data)
	}
}

// TODO: Handle error
func jsonError(rw web.ResponseWriter, err error) {
	if err != nil {
		jData, _ := json.Marshal(jsn{
			"Error": err.Error(),
		})
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(jData)
	}
}
