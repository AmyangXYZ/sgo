package main

import (
	"github.com/AmyangXYZ/sweetygo"
)

func main() {
	app := sweetygo.New()

	app.USE(sweetygo.Logger())
	app.Static("/static", "/home/amyang/Projects/SweetyGo/example/static")
	app.GET("/", home)
	app.GET("/api", api)
	app.GET("/usr/:user/:sex/:age", hello)

	app.RunServer(":16311")
}

func home(ctx *sweetygo.Context) {
	ctx.HTML(200, "Welcome\n")
}

func api(ctx *sweetygo.Context) {
	ctx.JSON(200, map[string]int{"uid": 001})
}

func hello(ctx *sweetygo.Context) {
	params := ctx.Params()
	body := "Hello" + params["usr"][0]
	ctx.HTML(200, body)
}
