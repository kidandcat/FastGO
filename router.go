package main

import "github.com/gocraft/web"

func setRouter() *web.Router {
	router := web.New(GlobalContext{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(web.StaticMiddleware("./static", web.StaticOption{})).
		NotFound((*GlobalContext).NotFound)

	loadControllers(router)

	return router
}

func loadControllers(router *web.Router) {
	UserController(router)
}
