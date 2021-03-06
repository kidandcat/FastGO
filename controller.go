package main

import (
	"encoding/json"
	"math/rand"
	"strings"

	redis "gopkg.in/redis.v5"

	"github.com/gocraft/web"
)

var db = map[string]dbc{}

// TODO: Accept storage parameter (memory, redis, redisCluster, postgre)
func Controller(router *web.Router, serviceName string, storage string, options jsn) {
	router.Subrouter(cxt, config.ApiPrefix+"/"+serviceName).
		//Find
		Get("/", (*GlobalContext).Find).
		//Get
		Get("/:id", (*GlobalContext).Get).
		//Create
		Post("/", (*GlobalContext).Create).
		//Update
		Put("/:id", (*GlobalContext).Update).
		//Patch
		Patch("/:id", (*GlobalContext).Patch).
		//Remove
		Delete("/:id", (*GlobalContext).Remove)

	switch storage {
	case "redis":
		db[serviceName] = Jredis(&redis.Options{
			Addr:     options["Addr"].(string),
			Password: options["Password"].(string),
			DB:       0,
		})
	default:
		db[serviceName] = Jmemory()
	}
}

// TODO: Generic database wrapper
// TODO: Accept ID parameter as queryString

// GET /messages
func (c *GlobalContext) Find(rw web.ResponseWriter, req *web.Request) {
	ser := strings.Split(req.RoutePath(), "/")
	req.Body.Close()
	service := ser[len(ser)-2]
	// TODO: Allow filters
	res := db[service].Find("*")
	jsonError(rw, res.Error)
	jsonAnswer(rw, res.Data)
}

// GET /messages/<id>
func (c *GlobalContext) Get(rw web.ResponseWriter, req *web.Request) {
	ser := strings.Split(req.RoutePath(), "/")
	id := req.PathParams["id"]
	req.Body.Close()
	service := ser[len(ser)-2]
	res := db[service].Get(id)
	jsonError(rw, res.Error)
	jsonAnswer(rw, jsonParse(res.Data))
}

// POST /messages
func (c *GlobalContext) Create(rw web.ResponseWriter, req *web.Request) {
	ser := strings.Split(req.RoutePath(), "/")
	service := ser[len(ser)-2]
	// TODO: Allow set Key as parameter at creating
	key := randomKey(service)
	//Parse data
	u, err1 := jsonReqParse(req, new(Anon))
	req.Body.Close()
	jsonError(rw, err1)
	res1, err2 := json.Marshal(u)
	jsonError(rw, err2)

	res2 := db[service].Set(key, string(res1))
	jsonError(rw, res2.Error)

	jsonAnswer(rw, jsn{
		"status": jsonParse(string(res1)),
		"key":    key,
	})
}

// PUT /messages[/<id>]
func (c *GlobalContext) Update(rw web.ResponseWriter, req *web.Request) {
	ser := strings.Split(req.RoutePath(), "/")
	service := ser[len(ser)-2]
	u, err1 := jsonReqParse(req, new(Anon))
	jsonError(rw, err1)
	r, err2 := json.Marshal(u)
	jsonError(rw, err2)

	res := db[service].Set(req.PathParams["id"], string(r))
	req.Body.Close()
	jsonError(rw, res.Error)
	jsonAnswer(rw, jsn{
		"status": res.Data,
	})
}

// PATCH /messages[/<id>]
func (c *GlobalContext) Patch(rw web.ResponseWriter, req *web.Request) {
	ser := strings.Split(req.RoutePath(), "/")
	service := ser[len(ser)-2]
	var uu map[string]string
	r1, err1 := jsonReqParse(req, uu)
	jsonError(rw, err1)
	r2 := db[service].Get(req.PathParams["id"])
	jsonError(rw, r2.Error)
	var rr map[string]interface{}
	err5 := json.Unmarshal([]byte(r2.Data), &rr)
	jsonError(rw, err5)

	for k, v := range r1.(map[string]interface{}) {
		rr[k] = v
	}

	res, err2 := json.Marshal(rr)
	jsonError(rw, err2)
	o := db[service].Set(req.PathParams["id"], string(res))
	req.Body.Close()
	jsonError(rw, o.Error)
	jsonAnswer(rw, jsn{
		"status": o.Data,
	})
}

// DELETE /messages[/<id>]
func (c *GlobalContext) Remove(rw web.ResponseWriter, req *web.Request) {
	ser := strings.Split(req.RoutePath(), "/")
	service := ser[len(ser)-2]
	r := db[service].Del(req.PathParams["id"])
	req.Body.Close()
	jsonError(rw, r.Error)
	jsonAnswer(rw, jsn{
		"status": r.Data,
	})
}

// TODO: Pass Key size as config parameter
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomKey(service string) string {
	b := make([]rune, 20)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	key := string(b)

	r := db[service].Exists(key)
	if r.Data == false && r.Error != nil {
		return key
	} else {
		return randomKey(service)
	}
}
