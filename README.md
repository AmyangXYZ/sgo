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
	"fmt"

	"github.com/AmyangXYZ/sweetygo"
)

func main() {
	app := sweetygo.New()

	app.USE(sweetygo.Logger())
	// app.GET("/static/*files", staticServer)
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
```

![example](https://raw.githubusercontent.com/AmyangXYZ/sweetygo/master/example/example.png)

## TODOs

- [ ] Context
- [ ] Unit Tests
- [ ] Built-in StaticServer
