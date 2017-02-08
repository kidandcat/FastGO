package main

import (
	"time"

	redis "gopkg.in/redis.v5"
)

//var MEMs = []container{}

type container struct {
	Memory map[string]string
}

type Res struct {
	res []string
}

func memory() *container {
	c := container{
		Memory: map[string]string{},
	}
	return &c
}

func (me *container) Keys(filter string) *redis.StringSliceCmd {
	keys := make([]string, 0, len((*me).Memory))
	for k, _ := range me.Memory {
		keys = append(keys, k)
	}
	return redis.NewStringSliceResult(keys, nil)
}

func (me *container) Set(key string, value interface{}, t time.Duration) *redis.StatusCmd {
	me.Memory[key] = value.(string)
	return redis.NewStatusResult(value.(string), nil)
}

func (me *container) Get(key string) *redis.StringCmd {
	return redis.NewStringResult(me.Memory[key], nil)
}

func (me *container) Del(key ...string) *redis.IntCmd {
	for _, i := range key {
		delete(me.Memory, i)
	}
	return redis.NewIntResult(1, nil)
}

func (me *container) Exists(key string) *redis.BoolCmd {
	if _, ok := me.Memory[key]; ok {
		return redis.NewBoolResult(true, nil)
	}
	return redis.NewBoolResult(false, nil)
}
