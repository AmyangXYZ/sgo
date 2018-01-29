# SweetyGo

SweetyGo is a simple, light and fast Web framework written in Go. 

## Features

- Pretty and fast router - based on radix tree
- Middleware Support
- Friendly to REST API
- No regexp or reflect
- Inspired by many excellent Go Web framework

## Installation

`go get github.com/AmyangXYZ/sweetygo`

## Example

```go
package main

import (
	"github.com/AmyangXYZ/sweetygo"
)

func main() {
	app := sweetygo.New()

	app.USE(sweetygo.Logger())
	app.Static("/static", "/home/amyang/Projects/SweetyGo/example/static")
	app.GET("/", home)
	app.POST("/api", api)
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

```

![example](https://raw.githubusercontent.com/AmyangXYZ/sweetygo/master/example/example.png)

## TODOs

- [ ] Session
- [ ] Render
- [ ] Some built-in Security Middleware
- [ ] Unit Tests
