# SweetyGo

SweetyGo is a simple, light and fast Web framework written in Go. 

The source is easy to learn, then you can make your own Go Web Framework!

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
	"html/template"
	"os"
	"time"

	"github.com/AmyangXYZ/sweetygo"
	"github.com/AmyangXYZ/sweetygo/middlewares"
	"github.com/dgrijalva/jwt-go"
)

var (
	rootDir        = "/home/amyang/Projects/SweetyGo/example"
	secretKey      = "CuteSweetie"
	requiredJWTMap = map[string]string{
		"/api":   "!GET",
		"/login": "GET",
		"/usr/*": "POST",
		"/api/2": "ALL",
	}
	listenPort = ":16311"
)

func main() {
	app := sweetygo.New(rootDir, template.FuncMap{})
	app.USE(middlewares.Logger(os.Stdout))
	app.USE(middlewares.JWT("Header", secretKey, requiredJWTMap))
	app.GET("/", home)
	app.GET("/api", biu)
	app.POST("/api", hello)
	app.POST("/login", login)
	app.GET("/usr/:user", usr)
	app.RunServer(listenPort)
}

func home(ctx *sweetygo.Context) {
	ctx.Set("baby", "Sweetie")
	ctx.Render(200, "index")
}

func biu(ctx *sweetygo.Context) {
	ctx.Text(200, "biu")
}

func login(ctx *sweetygo.Context) {
	usr := ctx.Param("usr")
	pwd := ctx.Param("pwd")
	if usr == "Amyang" && pwd == "biu" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Amyang"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		t, _ := token.SignedString([]byte(secretKey))
		ctx.JSON(200, 1, "success", map[string]string{"SG_Token": t})
		return
	}
	ctx.JSON(200, 0, "username or password error", nil)
}

func api(ctx *sweetygo.Context) {
	ctx.JSON(200, 1, "success", map[string]string{"version": 1.1})
}

func hello(ctx *sweetygo.Context) {
	usr := ctx.Get("userInfo").(*jwt.Token)
	claims := usr.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	ctx.Text(200, "Hello "+name)
}

func usr(ctx *sweetygo.Context) {
	ctx.Text(200, "Welcome home, "+ctx.Param("user"))
}


```

![example](https://raw.githubusercontent.com/AmyangXYZ/sweetygo/master/example/example.png)

My [Blog](https://amyang.xyz) is also powered by SweetyGo.

## TODOs

- [ ] More built-in Security Middleware
- [ ] Unit Tests
