# SGo

SGo is a simple, light and fast Web framework written in Go. 

The source is easy to learn, then you can make your own Go Web Framework!

## Features

- Pretty and fast router - based on radix tree
- Middleware Support
- Friendly to REST API
- No regexp or reflect
- QUIC Support
- Inspired by many excellent Go Web framework

## Installation

`go get github.com/AmyangXYZ/sgo`

## Example

### Simple

```go
package main

import (
    "github.com/AmyangXYZ/sgo"
)

func main() {
    app := sgo.New()
    app.GET("/", func(ctx *sgo.Context) error {
        return ctx.Text(200, "Hello Sweetie")
    })
    app.Run(":16311")
}

```

### Further

```go
package main

import (
    "html/template"
    "net/http"
    "os"
    "time"

    "github.com/AmyangXYZ/sgo"
    "github.com/AmyangXYZ/sgo/middlewares"
    "github.com/dgrijalva/jwt-go"
)

var (
    tplDir     = "templates"
    listenPort = ":16311"
    secretKey  = "CuteSweetie"
    jwtSkipper = func(ctx *sgo.Context) bool {
        if ctx.Path() == "/" ||
            (ctx.Path() == "/api" && ctx.Method() == "GET") ||
            (ctx.Path() == "/login" && ctx.Method() == "POST") ||
            (len(ctx.Path()) > 8 && ctx.Path()[0:8] == "/static/") {
            return true
        }
        return false
    }
)

func main() {
    app := sgo.New()
    app.SetTemplates(tplDir, template.FuncMap{})

    app.USE(middlewares.Logger(os.Stdout, middlewares.DefaultSkipper))
    app.USE(middlewares.JWT("Header", secretKey, jwtSkipper))

    app.GET("/", home)
    app.GET("/static/*files", static)
    app.GET("/api", biu)
    app.POST("/api", hello)
    app.POST("/login", login)
    app.GET("/usr/:user", usr)

    app.Run(listenPort)

    // Or use QUIC
    // app.RunOverQUIC(listenPort, "fullchain.pem", "privkey.pem")
}

func home(ctx *sgo.Context) {
    ctx.Set("baby", "Sweetie")
    ctx.Render(200, "index")
}

func static(ctx *sgo.Context) {
    staticHandle := http.StripPrefix("/static",
        http.FileServer(http.Dir("./static")))
    staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
}

func biu(ctx *sgo.Context) {
    ctx.Text(200, "biu")
}

func login(ctx *sgo.Context) {
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

func api(ctx *sgo.Context) {
    ctx.JSON(200, 1, "success", map[string]string{"version": "1.1"})
}

func hello(ctx *sgo.Context) {
    usr := ctx.Get("userInfo").(*jwt.Token)
    claims := usr.Claims.(jwt.MapClaims)
    name := claims["name"].(string)
    ctx.Text(200, "Hello "+name)
}

func usr(ctx *sgo.Context) {
    ctx.Text(200, "Welcome home, "+ctx.Param("user"))
}

```

![example](https://raw.githubusercontent.com/AmyangXYZ/sgo/master/example/example.png)

My [Blog](https://amyang.xyz) is also powered by SGo.

## License

MIT