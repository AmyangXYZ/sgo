package main

import (
	"fmt"

	"github.com/AmyangXYZ/sweetygo"
)

func main() {
	app := sweetygo.New()

	app.USE(sweetygo.Logger())
	app.Static("/static", "/home/amyang/Projects/SweetyGo/example/static")
	app.GET("/", home)
	app.POST("/api", home)
	app.GET("/usr/:user/:sex/:age", hello)

	app.RunServer(":16311")
}

func home(ctx *sweetygo.Context) {
	ctx.Resp.WriteHeader(200)
	fmt.Fprintf(ctx.Resp, "Welcome \n")
}

func hello(ctx *sweetygo.Context) {
	params := ctx.Params()
	ctx.Resp.WriteHeader(200)
	fmt.Fprintf(ctx.Resp, "Hello %s\n", params["user"][0])
}
