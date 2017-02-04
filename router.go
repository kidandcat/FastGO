package main

import (
	"fmt"

	"github.com/gocraft/web"
)

func setRouter() *web.Router {
	router := web.New(GlobalContext{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(web.StaticMiddleware("./"+config.StaticFolder, web.StaticOption{})).
		NotFound((*GlobalContext).NotFound)

	loadControllers(router)

	return router
}

func loadControllers(router *web.Router) {
	UserController(router)
}

func (c *GlobalContext) NotFound(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "Route not found, sorry")
}
