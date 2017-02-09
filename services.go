package main

import "github.com/gocraft/web"

func Services(router *web.Router) {

	//Controller(router, "memory", "memory", jsn{})
	Controller(router, "ticket", "redis", jsn{
		"Addr":     "localhost:6667",
		"Password": "",
	})
	//Controller(router, "ticket", "memory", jsn{})
}
