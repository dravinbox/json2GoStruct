package main

import (
	"flag"
	. "json2GoStruct/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

var (
	json = flag.String("json", "{}", "input json string")
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("view", ".html"))
	app.Get("/", func(context context.Context) {
		//context.JSON(Address{"guangdong","guangdong"})
		context.View("index.html" )
	})
	app.Post("/convert_form_action",ConvertAction)
	app.Run(iris.Addr(":8080"))

}
