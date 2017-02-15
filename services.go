package main

import "github.com/gocraft/web"

func Services(router *web.Router) {

	//Controller(router, "user", "memory", jsn{})
	// Controller(router, "ticket", "redis", jsn{
	// 	"Addr":     "localhost:6666",
	// 	"Password": "",
	// })
	Controller(router, "user", "memory", jsn{})
}
