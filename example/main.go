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
