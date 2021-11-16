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
        return ctx.Text(200, "Hello")
    })
    app.Run(":16311")
}

```

### Further

For vue2 projects, add `module.exports = {assetsDir: 'static', css: { extract: false }}` to vue.config.js, then `npm run build && tar caf dist.tar.xz dist` and copy dist.tar.xz and run `./deployFrontend.sh`.

```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AmyangXYZ/sgo"
	"github.com/gorilla/websocket"
)

const (
	addr = ":8888"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	app := sgo.New()
	app.SetTemplates("./templates", nil)
	app.GET("/", index)
	app.GET("/static/*files", static)

	app.GET("/api/boottime", getBootTime)
	app.GET("/ws/comm", wsComm)
	app.POST("/api/link/:name", postHandler)
	app.OPTIONS("/api/link/:name", sgo.PreflightHandler)

	if err := app.Run(addr); err != nil {
		log.Fatal("Listen error", err)
	}
}

// Index page handler.
func index(ctx *sgo.Context) error {
	return ctx.Render(200, "index")
}

// Static files handler.
func static(ctx *sgo.Context) error {
	staticHandle := http.StripPrefix("/static",
		http.FileServer(http.Dir("./static")))
	staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
	return nil
}

func getBootTime(ctx *sgo.Context) error {
	return ctx.Text(200, fmt.Sprintf("%d", 20))
}

func wsComm(ctx *sgo.Context) error {
	ws, err := upgrader.Upgrade(ctx.Resp, ctx.Req, nil)
	breakSig := make(chan bool)
	if err != nil {
		return err
	}
	fmt.Println("ws/comm connected")
	defer func() {
		ws.Close()
		fmt.Println("ws/comm client closed")
	}()
	go func() {
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				breakSig <- true
			}
		}
	}()
	for {
		select {
		// case l := <-LogsComm:
		// 	ws.WriteJSON(l)
		case <-breakSig:
			return errors.New("stop ws")
		}
	}
}

func postHandler(ctx *sgo.Context) error {
	// param request
	fmt.Println(ctx.Params)
	// json body request
	body, err := ioutil.ReadAll(ctx.Req.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	return ctx.Text(200, "xx")
}
```

![example](https://raw.githubusercontent.com/AmyangXYZ/sgo/master/example/example.png)

My [Blog](https://amyang.xyz) is also powered by SGo.

## License

MIT