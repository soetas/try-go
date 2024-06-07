package json

import "encoding/json"

func Stringify(value interface{}) string {
	json, _ := json.Marshal(value)
	return string(json)
}
