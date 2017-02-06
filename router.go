package main

import (
	"fmt"

	"github.com/gocraft/web"
)

var cxt GlobalContext

func setRouter() *web.Router {
	cxt = GlobalContext{}
	router := web.New(cxt).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(web.StaticMiddleware("./"+config.StaticFolder, web.StaticOption{})).
		NotFound((cxt).NotFound)

	loadControllers(router)

	return router
}

func loadControllers(router *web.Router) {
	UserController(router)
}

func (c *GlobalContext) NotFound(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "Route not found, sorry")
}
