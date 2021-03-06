package controller

import (
	"fmt"
	. "github.com/dravinbox/json2GoStruct/model"
	. "github.com/dravinbox/json2GoStruct/util"
	"github.com/kataras/iris"
	"reflect"
	"strings"
)

func ConvertAction(ctx iris.Context)  {
	jsondata := JsonData{}
	err := ctx.ReadForm(&jsondata)
	if err !=nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}

	jsondata.GoStruct = Json2StructString(jsondata.JsonStr,jsondata.MainName)

	ctx.View("index.html",jsondata)

}




/**
把json 字符串转化为go struct 的字符串
 */
func Json2StructString(json string,structName string) (structString string){
	structName = StrFirstToUpper(structName)
	jsonMap, e := Json2Map(json)

	if e != nil {
		fmt.Println(e)
		return
	}

	var goStruct []string
	Convert(jsonMap,structName,&goStruct)
	structString = strings.Join(goStruct,"\n")
	return

}

/**
递归分析map中是否含有(map[string]interface{},[]interface{})类型，有就转化为JsonObject
*/
func Convert(jsonMap map[string]interface{}, structName string,goStruct *[]string) {
	// 遍历map
	for k, v := range jsonMap {
		switch v.(type) {

		//如果是json对象
		case map[string]interface{}:
			v, _ := v.(map[string]interface{})
			Convert(v, StrFirstToUpper(k),goStruct)
			jsonMap[k] = JsonObject{Name: StrFirstToUpper(k)}
		//如果是json数组
		case []interface{}:

			v, _ := v.([]interface{})
			length := len(v)
			var t map[string]interface{}
			if length > 0 {
				t, _ = v[0].(map[string]interface{})

			}
			Convert(t, StrFirstToUpper(k),goStruct)
			jsonMap[k] = JsonObject{Name: StrFirstToUpper(k)}

		}

	}
	Print(jsonMap, structName,goStruct)

}

/**
把只含有以下类型(string,float64,JsonObject)的map 转化为Go Struct
*/
func Print(jsonBaseMap map[string]interface{}, structName string,goStruct *[]string) {

	var stringSlice []string
	for k, v := range jsonBaseMap {
		valueType := reflect.TypeOf(v)

		jsonTip := fmt.Sprintf("`json:\"%s\"`", k)
		var line string
		if valueType == reflect.TypeOf(JsonObject{}) {
			line = fmt.Sprintf("\t%s %s %s", StrFirstToUpper(k), v, jsonTip)

		} else {
			line = fmt.Sprintf("\t%s %s %s", StrFirstToUpper(k), valueType, jsonTip)

		}
		stringSlice = append(stringSlice, line)
	}
	content := strings.Join(stringSlice, "\n")
	result := fmt.Sprintf("type %s struct {\n%s\n}", structName, content)
	*goStruct = append(*goStruct, result)
	//fmt.Println(result)

}
