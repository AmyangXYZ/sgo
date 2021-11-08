package main

import (
	"github.com/AmyangXYZ/sgo"
)

func main() {
	app := sgo.New()
	app.GET("/", func(ctx *sgo.Context) error {
		return ctx.Text(200, "Hello biu")
	})
	app.Run(":16311")
}
