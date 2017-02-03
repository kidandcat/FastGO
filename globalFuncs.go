package main

import (
	"fmt"

	"github.com/gocraft/web"
)

func (c *GlobalContext) NotFound(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "Route not found, sorry")
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
