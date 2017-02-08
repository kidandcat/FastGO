package main

import redis "gopkg.in/redis.v5"

func Jredis(options *redis.Options) *redisStorage {
	c := redisStorage{
		Memory: redis.NewClient(options),
	}
	return &c
}

func (me *redisStorage) Find(filter string) ResultJsnArray {
	var us []jsn

	keys, err := me.Memory.Keys("*").Result()

	for k, v := range keys {
		us = append(us, jsn{
			"Data": jsonParse(me.Get(v).Data),
			"Key":  k,
		})
	}
	return ResultJsnArray{
		Data:  us,
		Error: err,
	}
}

func (me *redisStorage) Set(key string, value string) ResultBoolean {
	err := me.Memory.Set(key, value, 0).Err()
	return ResultBoolean{
		Data:  true,
		Error: err,
	}
}

func (me *redisStorage) Get(key string) ResultString {
	r, err := me.Memory.Get(key).Result()
	res := ResultString{
		Data:  r,
		Error: err,
	}
	return res
}

func (me *redisStorage) Del(key string) ResultBoolean {
	err := me.Memory.Del(key).Err()
	return ResultBoolean{
		Data:  true,
		Error: err,
	}
}

func (me *redisStorage) Exists(key string) ResultBoolean {
	r, err := me.Memory.Exists(key).Result()

	return ResultBoolean{
		Data:  r,
		Error: err,
	}
}
