package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocraft/web"
	"gopkg.in/redis.v5"
)

type User struct {
	Name string `json:"name"`
}

var db *redis.Client

func UserController(router *web.Router) {
	router.Subrouter(cxt, "/user").
		Get("/", (*GlobalContext).Find).
		Get("/:id", (*GlobalContext).Get).
		Post("/", (*GlobalContext).Create).
		Put("/:id", (*GlobalContext).Update).
		Patch("/:id", (*GlobalContext).Patch).
		Delete("/:id", (*GlobalContext).Remove)

	db = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,
	})
}

// GET /messages
func (c *GlobalContext) Find(rw web.ResponseWriter, req *web.Request) {
	keys, er := db.Keys("*").Result()
	jsonError(rw, er)
	var us []User
	for _, e := range keys {
		val, _ := db.Get(e).Result()
		var u User
		_ = json.Unmarshal([]byte(val), &u)
		us = append(us, u)
	}
	jsonAnswer(rw, us)
}

// GET /messages/<id>
func (c *GlobalContext) Get(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "GET executed")
}

// POST /messages
func (c *GlobalContext) Create(rw web.ResponseWriter, req *web.Request) {
	u, err1 := jsonParse(req, new(User))
	jsonError(rw, err1)
	res, err2 := json.Marshal(u)
	jsonError(rw, err2)
	err3 := db.Set("key", res, 0).Err()
	jsonError(rw, err3)
	jsonAnswer(rw, jsn{
		"status": "OK",
	})
}

// PUT /messages[/<id>]
func (c *GlobalContext) Update(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "UPDATE executed")
}

// PATCH /messages[/<id>]
func (c *GlobalContext) Patch(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "PATCH executed")
}

// DELETE /messages[/<id>]
func (c *GlobalContext) Remove(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "REMOVE executed")
}
