package main

type ResultString struct {
	Data  string
	Error error
}
type ResultStringArray struct {
	Data  []string
	Error error
}
type ResultJsnArray struct {
	Data  []jsn
	Error error
}
type ResultJsn struct {
	Data  []string
	Error error
}
type ResultBoolean struct {
	Data  bool
	Error error
}

type DataKey struct {
	Key  string `json:"key"`
	Data Anon   `json:"data"`
}

type Anon interface{}

// TODO: Generic database wrapper
type dbc interface {
	Find(string) ResultJsnArray
	Get(string) ResultString
	Set(string, string) ResultBoolean
	Del(string) ResultBoolean
	Exists(string) ResultBoolean
}

type Config struct {
	Port         string `json:"port"`
	StaticFolder string `json:"staticFolder"`
}

type GlobalContext struct{}

type jsn map[string]interface{}

type storage struct {
	Memory map[string]string
}
