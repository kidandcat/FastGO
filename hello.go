package main

import (
	"fmt"
	"net/http"
)

var port string = "8887"

type GlobalContext struct {
	Name string
}

func main() {
	router := setRouter()

	fmt.Println("Server listening at ", port)
	panicOnError(http.ListenAndServe("localhost:"+port, router))
}
