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
	"time"

	"github.com/AmyangXYZ/sweetygo"
	"github.com/AmyangXYZ/sweetygo/middlewares"
	"github.com/dgrijalva/jwt-go"
)

var (
	secretKey      = "CuteSweetie"
	requiredJWTMap = map[string]string{
		"/api":   "!GET",
		"/login": "GET",
		"/usr/*": "POST",
		"/api/2": "ALL",
	}
)

func main() {
	rootDir := "/home/amyang/Projects/SweetyGo/example"
	app := sweetygo.New(rootDir)
	app.USE(middlewares.Logger())
	app.USE(middlewares.JWT(secretKey, requiredJWTMap))
	app.GET("/", home)
	app.GET("/api", biu)
	app.POST("/api", hello)
	app.POST("/login", login)
	app.GET("/usr/:user", usr)

	app.RunServer(":16311")
}

func home(ctx *sweetygo.Context) {
	ctx.Set("baby", "Sweetie")
	ctx.Render(200, "index")
}

func biu(ctx *sweetygo.Context) {
	ctx.Text(200, "biu")
}

func login(ctx *sweetygo.Context) {
	params := ctx.ParseForm()
	usr := params["usr"][0]
	pwd := params["pwd"][0]
	if usr == "Amyang" && pwd == "biu" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Amyang"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		t, _ := token.SignedString([]byte(secretKey))
		ctx.SetCookie("sgtoken", t)
		ctx.JSON(200, map[string]string{"token": t})
	}
}

func api(ctx *sweetygo.Context) {
	ctx.JSON(200, map[string]int{"uid": 001})
}

func hello(ctx *sweetygo.Context) {
	usr := ctx.Get("user").(*jwt.Token)
	claims := usr.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	ctx.Text(200, "Hello "+name)
}

func usr(ctx *sweetygo.Context) {
	params := ctx.ParseForm()
	ctx.Text(200, "Welcome home, "+params["user"][0])
}


```

![example](https://raw.githubusercontent.com/AmyangXYZ/sweetygo/master/example/example.png)

## TODOs

- [ ] Some built-in Security Middleware
- [ ] Unit Tests
