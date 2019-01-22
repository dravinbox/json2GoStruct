package main

import (
	"flag"
	"fmt"
	"github.com/dravinbox/json2GoStruct/controller"
	"github.com/dravinbox/json2GoStruct/util"
)

var (
 json = flag.String("json","{}","input json string")

)





func main() {

	//var str ="{\"user_name\":\"zry\",\"age\":18,\"profile\":{\"color\":\"red\",\"size\":12},\"tel\":[],\"address\":[{\"city\":\"Guangzhou\",\"province\":\"Guangdong\"}]}"
	flag.Parse()

	//fmt.Println(*json)
	jsonMap, e := util.Json2Map(*json)
	//jsonMap, e := util.Json2Map(str)

	if e !=nil {
		fmt.Println(e)
		return
	}

	controller.Convert(jsonMap,"A")

}


