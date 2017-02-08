package main

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	"github.com/gocraft/web"
	"gopkg.in/redis.v5"
)

type Anon interface{}

type DataKey struct {
	Key  string `json:"key"`
	Data Anon   `json:"data"`
}

var db = map[string]dbc{}

// TODO: Generic database wrapper
type dbc interface {
	Get(string) *redis.StringCmd
	Set(string, interface{}, time.Duration) *redis.StatusCmd
	Del(...string) *redis.IntCmd
	Keys(string) *redis.StringSliceCmd
	Exists(string) *redis.BoolCmd
}

// TODO: Accept storage parameter (memory, redis, redisCluster, postgre)
func Controller(router *web.Router, serviceName string, addr string, pass string) {
	router.Subrouter(cxt, "/"+serviceName).
		Get("/", (*GlobalContext).Find).
		Get("/:id", (*GlobalContext).Get).
		Post("/", (*GlobalContext).Create).
		Put("/:id", (*GlobalContext).Update).
		Patch("/:id", (*GlobalContext).Patch).
		Delete("/:id", (*GlobalContext).Remove)

		// TODO: Switch storage parameter
	if addr == "" {
		db[serviceName] = memory()
	} else {
		db[serviceName] = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pass,
			DB:       0,
		})
	}
}

// TODO: Generic database wrapper
// TODO: Accept ID parameter as queryString

// GET /messages
func (c *GlobalContext) Find(rw web.ResponseWriter, req *web.Request) {
	service := strings.Split(req.RoutePath(), "/")[1]
	// TODO: Allow filters
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
	// TODO: Allow set Key as parameter at creating
	key := randomKey(service)
	u, err1 := jsonParse(req, new(Anon))
	jsonError(rw, err1)
	res, err2 := json.Marshal(u)
	jsonError(rw, err2)
	err3 := db[service].Set(key, string(res), 0).Err()
	jsonError(rw, err3)
	jsonAnswer(rw, jsn{
		"status": "OK",
		"key":    key,
	})
}

// PUT /messages[/<id>]
func (c *GlobalContext) Update(rw web.ResponseWriter, req *web.Request) {
	service := strings.Split(req.RoutePath(), "/")[1]
	u, err1 := jsonParse(req, new(Anon))
	jsonError(rw, err1)
	res, err2 := json.Marshal(u)
	jsonError(rw, err2)
	err3 := db[service].Set(req.PathParams["id"], string(res), 0).Err()
	jsonError(rw, err3)
	jsonAnswer(rw, u)
}

// PATCH /messages[/<id>]
func (c *GlobalContext) Patch(rw web.ResponseWriter, req *web.Request) {
	service := strings.Split(req.RoutePath(), "/")[1]
	var uu map[string]string
	u, err1 := jsonParse(req, uu)
	jsonError(rw, err1)
	r, err4 := db[service].Get(req.PathParams["id"]).Result()
	jsonError(rw, err4)
	var rr map[string]string
	err5 := json.Unmarshal([]byte(r), rr)
	jsonError(rw, err5)

	for k, v := range uu {
		rr[k] = v
	}

	res, err2 := json.Marshal(rr)
	jsonError(rw, err2)
	err3 := db[service].Set(req.PathParams["id"], string(res), 0).Err()
	jsonError(rw, err3)
	jsonAnswer(rw, u)
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

// TODO: Pass Key size as config parameter
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomKey(service string) string {
	b := make([]rune, 20)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	key := string(b)

	r, _ := db[service].Exists(key).Result()
	if r == false {
		return key
	} else {
		return randomKey(service)
	}
}
