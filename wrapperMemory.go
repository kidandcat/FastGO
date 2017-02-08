package main

import "errors"

func memory() *storage {
	c := storage{
		Memory: map[string]string{},
	}
	return &c
}

func (me *storage) Find(filter string) ResultJsnArray {
	var us []jsn
	//keys := make([]string, 0, len((*me).Memory))
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

//
// func (me *storage) Keys(filter string) *redis.StringSliceCmd {
// 	keys := make([]string, 0, len((*me).Memory))
// 	for k, _ := range me.Memory {
// 		keys = append(keys, k)
// 	}
// 	return redis.NewStringSliceResult(keys, nil)
// }

func (me *storage) Set(key string, value string) ResultBoolean {
	me.Memory[key] = value
	return ResultBoolean{
		Data:  true,
		Error: nil,
	}
}

func (me *storage) Get(key string) ResultString {
	res := ResultString{
		Data:  me.Memory[key],
		Error: nil,
	}
	return res
}

func (me *storage) Del(key string) ResultBoolean {
	delete(me.Memory, key)
	return ResultBoolean{
		Data:  true,
		Error: nil,
	}
}

func (me *storage) Exists(key string) ResultBoolean {
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
