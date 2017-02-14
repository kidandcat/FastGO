package main

import "errors"

func Jmemory() *memoryStorage {
	c := memoryStorage{
		Memory: map[string]string{},
	}
	return &c
}

func (me *memoryStorage) Find(filter string) ResultJsnArray {
	var us []jsn
	for k, v := range me.Memory {
		us = append(us, jsn{
			"Data": jsonParse(v),
			"Key":  k,
		})
	}
	return ResultJsnArray{
		Data:  us,
		Error: nil,
	}
}

func (me *memoryStorage) Set(key string, value string) ResultBoolean {
	me.Memory[key] = value
	return ResultBoolean{
		Data:  true,
		Error: nil,
	}
}

func (me *memoryStorage) Get(key string) ResultString {
	res := ResultString{
		Data:  me.Memory[key],
		Error: nil,
	}
	return res
}

func (me *memoryStorage) Del(key string) ResultBoolean {
	delete(me.Memory, key)
	return ResultBoolean{
		Data:  true,
		Error: nil,
	}
}

func (me *memoryStorage) Exists(key string) ResultBoolean {
	if _, ok := me.Memory[key]; ok {
		return ResultBoolean{
			Data:  true,
			Error: nil,
		}
	}
	return ResultBoolean{
		Data:  false,
		Error: errors.New("Not Found"),
	}
}
