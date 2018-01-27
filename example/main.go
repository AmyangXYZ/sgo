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

func home(c *sweetygo.Context) {
	c.Resp.WriteHeader(200)
	fmt.Fprintf(c.Resp, "Welcome \n")
}

func hello(c *sweetygo.Context) {
	params := c.Params()
	c.Resp.WriteHeader(200)
	fmt.Fprintf(c.Resp, "Hello %s\n", params["user"][0])
}
