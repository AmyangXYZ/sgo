package main

import (
	"html/template"
	"net/http"

	"github.com/AmyangXYZ/sgo"
)

var (
	tplDir     = "templates"
	listenPort = "amyang.xyz:443"
)

func main() {
	app := sgo.New()
	app.SetTemplates(tplDir, template.FuncMap{})

	app.Any("/", home)
	app.GET("/static/*files", static)
	app.GET("/api", biu)
	app.GET("/usr/:user", usr)

	app.Run(listenPort)
	// Or use QUIC
	// app.RunOverQUIC(listenPort, "fullchain.pem", "privkey.pem")
}

func home(ctx *sgo.Context) error {
	ctx.Set("biu", "S")
	return ctx.Render(200, "index")
}

func static(ctx *sgo.Context) error {
	staticHandle := http.StripPrefix("/static",
		http.FileServer(http.Dir("./static")))
	staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
	return nil
}

func biu(ctx *sgo.Context) error {
	return ctx.Text(200, "biu")
}

func api(ctx *sgo.Context) error {
	return ctx.JSON(200, 1, "success", map[string]string{"version": "1.1"})
}

func usr(ctx *sgo.Context) error {
	return ctx.Text(200, "Welcome home, "+ctx.Param("user"))
}
