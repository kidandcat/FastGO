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
		Middleware(web.StaticMiddleware("./"+config.StaticFolder, web.StaticOption{IndexFile: "index.html"})).
		NotFound((cxt).NotFound)

	Services(router)

	return router
}

func (c *GlobalContext) NotFound(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "404 Route not found, sorry")
}
