package main

import (
	"fmt"

	"github.com/gocraft/web"
)

type UserContext struct {
	*GlobalContext
	Name string
}

func UserController(router *web.Router) {
	router.Subrouter(UserContext{}, "/user").
		Get("/:name", (*UserContext).SayName).
		Get("/", (*UserContext).SayHello)
}

func (c *UserContext) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "Hello World!")
}

func (c *UserContext) SayName(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "Hello ", req.PathParams["name"], " World!")
}
