package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocraft/web"
	"gopkg.in/redis.v5"
)

type Anon interface{}

type DataKey struct {
	Key  string `json:"key"`
	Data Anon   `json:"data"`
}

var db map[string]*redis.Client

func Controller(router *web.Router, serviceName string, addr string, pass string) {
	router.Subrouter(cxt, "/"+serviceName).
		Get("/", (*GlobalContext).Find).
		Get("/:id", (*GlobalContext).Get).
		Post("/", (*GlobalContext).Create).
		Put("/:id", (*GlobalContext).Update).
		Patch("/:id", (*GlobalContext).Patch).
		Delete("/:id", (*GlobalContext).Remove)

	db = map[string]*redis.Client{}

	db[serviceName] = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})
}

// GET /messages
func (c *GlobalContext) Find(rw web.ResponseWriter, req *web.Request) {
	service := strings.Split(req.RoutePath(), "/")[1]
	keys, er := db[service].Keys("*").Result()
	jsonError(rw, er)
	var us []DataKey
	for _, e := range keys {
		val, _ := db[service].Get(e).Result()
		var u Anon
		_ = json.Unmarshal([]byte(val), &u)
		us = append(us, DataKey{
			Data: u,
			Key:  e,
		})
	}
	jsonAnswer(rw, us)
}

// GET /messages/<id>
func (c *GlobalContext) Get(rw web.ResponseWriter, req *web.Request) {
	service := strings.Split(req.RoutePath(), "/")[1]
	r, err1 := db[service].Get(req.PathParams["id"]).Result()
	jsonError(rw, err1)
	var u Anon
	_ = json.Unmarshal([]byte(r), &u)
	jsonAnswer(rw, u)
}

// POST /messages
func (c *GlobalContext) Create(rw web.ResponseWriter, req *web.Request) {
	service := strings.Split(req.RoutePath(), "/")[1]
	key, err := db[service].RandomKey().Result()
	jsonError(rw, err)
	u, err1 := jsonParse(req, new(Anon))
	jsonError(rw, err1)
	res, err2 := json.Marshal(u)
	jsonError(rw, err2)
	err3 := db[service].Set(key, res, 0).Err()
	jsonError(rw, err3)
	jsonAnswer(rw, jsn{
		"status": "OK",
		"key":    key,
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
	service := strings.Split(req.RoutePath(), "/")[1]
	r, err := db[service].Del(req.PathParams["id"]).Result()
	jsonError(rw, err)
	if r == 1 {
		jsonAnswer(rw, jsn{
			"status": "OK",
		})
	} else {
		jsonAnswer(rw, jsn{
			"Error": "Not Found",
		})
	}
}
