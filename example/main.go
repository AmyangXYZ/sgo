package main

import (
	"github.com/AmyangXYZ/sweetygo"
)

func main() {
	rootDir := "/home/amyang/Projects/SweetyGo/example"
	app := sweetygo.New(rootDir)
	app.USE(sweetygo.Logger())
	app.GET("/", home)
	app.POST("/api", api)
	app.GET("/usr/:user/:sex/:age", hello)

	app.RunServer(":16311")
}

func home(ctx *sweetygo.Context) {
	ctx.SetVar("user", "Sweetie")
	ctx.Render(200, "index")
}

func api(ctx *sweetygo.Context) {
	ctx.JSON(200, map[string]int{"uid": 001})
}

func hello(ctx *sweetygo.Context) {
	params := ctx.ParseForm()
	body := "Hello " + params["user"][0]
	ctx.HTML(200, body)
}
