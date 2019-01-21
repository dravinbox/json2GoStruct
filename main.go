package main

import (
	"fmt"
	"github.com/dravinbox/json2GoStruct/controller"
	"github.com/dravinbox/json2GoStruct/util"
)



func main() {

	var str ="{\"user_name\":\"zry\",\"age\":18,\"profile\":{\"color\":\"red\",\"size\":12},\"tel\":[],\"address\":[{\"city\":\"Guangzhou\",\"province\":\"Guangdong\"}]}"

	jsonMap, e := util.Json2Map(str)
	if e !=nil {
		fmt.Println(e)
		return
	}

	controller.Convert(jsonMap,"A")

}


