package util

import (
	"encoding/json"
	"github.com/json-iterator/go"
)

/**
json 转 map[string]interface{}
使用encoding/json
*/
func Json2MapWEcodingJson(str string) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil

}
/**
json 转 map[string]interface{}
使用jsoniter 性能较高
*/
func Json2Map(str string) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := jsoniter.Unmarshal([]byte(str), &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil

}
