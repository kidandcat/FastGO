package main

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

var port string = "8887"

type GlobalContext struct {
	Name string
}

func main() {
	router := web.New(GlobalContext{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(web.StaticMiddleware("./static", web.StaticOption{})).
		NotFound((*GlobalContext).NotFound)

	UserController(router)

	fmt.Println("Server listening at ", port)
	err := http.ListenAndServe("localhost:"+port, router)
	panicOnError(err)
}

func (c *GlobalContext) NotFound(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "Route not found, sorry")
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
