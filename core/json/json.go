package json

import "github.com/syariatifaris/arkeus/core/json/jsonq"

//Query contract
type Query interface {
	FindInt(fields ...string) (int, error)
	FindString(fields ...string) (string, error)
	FindBool(fields ...string) (bool, error)
	FindFloat64(fields ...string) (float64, error)
	FindObject(fields ...string) (interface{}, error)
}

//jsonQuery creates new jsonq
func jsonQuery(jsonStr string) Query {
	return jsonq.CreateJsonQuery(jsonStr)
}

//Searchable as string
type Searchable string

//FindInt finds an integer value in json from fields name and depth
func (s Searchable) FindInt(fields ...string) (int, error) {
	return jsonQuery(string(s)).FindInt(fields...)
}

//FindString finds a string value in json from fields name and depth
func (s Searchable) FindString(fields ...string) (string, error) {
	return jsonQuery(string(s)).FindString(fields...)
}

//FindBool finds a bool value in json from fields name and depth
func (s Searchable) FindBool(fields ...string) (bool, error) {
	return jsonQuery(string(s)).FindBool(fields...)
}

//FindObject finds a object value in json from fields name and depth
func (s Searchable) FindObject(fields ...string) (interface{}, error) {
	return jsonQuery(string(s)).FindObject(fields...)
}

//FindFloat64 finds a float64 value in json from fields name and depth
func (s Searchable) FindFloat64(fields ...string) (float64, error) {
	return jsonQuery(string(s)).FindFloat64(fields...)
}
