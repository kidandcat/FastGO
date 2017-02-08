package main

import "github.com/gocraft/web"

func Services(router *web.Router) {
	//
	//                            Redis Address
	//             endpoint name       |     Redis Password
	//                   v             v          v
	//Controller(router, "user1", "localhost:6666", "")
	//
	// If third parameter is empty, then in memory storage is used
	Controller(router, "user2", "", "")
}
