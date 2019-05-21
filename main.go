package main

import (
	"fmt"
	. "github.com/dravinbox/json2GoStruct/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	flag "github.com/spf13/pflag"
	"io/ioutil"
)

var (
	web  bool
	name string
	file string
)

func main() {
	flag.BoolVarP(&web, "web", "w", false, "start web interface")
	flag.StringVarP(&name, "name", "n", "Hello", "Go struct name")
	flag.StringVarP(&file, "file", "f", "./file.json", "path to the json file")

	flag.Parse()




	if web {
		startWeb()
	}else{
		cmdLine()
	}

}

func startWeb() {
	app := iris.New()
	app.RegisterView(iris.HTML("view", ".html"))
	app.Get("/", func(context context.Context) {
		//context.JSON(Address{"guangdong","guangdong"})
		context.View("index.html")
	})
	app.Post("/convert_form_action", ConvertAction)
	app.Run(iris.Addr(":8080"))

}

func cmdLine(){
	bytes, e := ioutil.ReadFile(file)
	if e!=nil{
		fmt.Println(e)
		return
	}

	json := string(bytes)
	structString := Json2StructString(json, name)
	fmt.Println(structString)
}