package main

import (
	"fmt"
	"net/http"

	"github.com/AmyangXYZ/sweetygo"
)

func main() {
	app := sweetygo.New(root)

	app.USE(sweetygo.Logger())
	app.GET("/static/*files", staticServer)
	app.GET("/", home)
	app.POST("/api", home)
	app.GET("/usr/:user/:sex/:age", hello)

	app.RunServer(":8080")
}

func root(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	http.NotFound(w, r)
}

func home(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Welcome \n")
}

func staticServer(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	staticHandle := http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static")))
	staticHandle.ServeHTTP(w, r)
}

func hello(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	r.ParseForm()
	params := r.Form
	w.WriteHeader(200)
	fmt.Fprintf(w, "Hello %s\n", params["user"][0])
}
