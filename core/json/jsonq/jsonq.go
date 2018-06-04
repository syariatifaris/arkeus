package jsonq

import (
	"encoding/json"
	"strings"

	"github.com/jmoiron/jsonq"
)

//CreateJsonQuery creates jsonq instance
func CreateJsonQuery(jsonStr string) *JsonQImpl {
	jsonStr = strings.Replace(jsonStr, `\"`, `"`, -1)
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	if jq == nil {
		return nil
	}

	return &JsonQImpl{
		jq: jq,
	}
}

//JsonQImpl structure implementation
type JsonQImpl struct {
	jq *jsonq.JsonQuery
}

//FindInt finds an integer value from json
func (j *JsonQImpl) FindInt(fields ...string) (int, error) {
	return j.jq.Int(fields...)
}

//FindString finds a string value from json
func (j *JsonQImpl) FindString(fields ...string) (string, error) {
	return j.jq.String(fields...)
}

//FindBool finds a bool value from json
func (j *JsonQImpl) FindBool(fields ...string) (bool, error) {
	return j.jq.Bool(fields...)
}

//FindFloat64 finds a float64 value from json
func (j *JsonQImpl) FindFloat64(fields ...string) (float64, error) {
	return j.jq.Float(fields...)
}

//FindObject finds a object value from json
func (j *JsonQImpl) FindObject(fields ...string) (interface{}, error) {
	return j.jq.Object(fields...)
}
