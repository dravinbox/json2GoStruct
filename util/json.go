package util

import (
	"encoding/json"
)

/**
json 转 map[string]interface{}
*/
func Json2Map(str string) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil

}
