package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/AmyangXYZ/sgo"
	"github.com/AmyangXYZ/sgo/middlewares"
)

var (
	tplDir        = "templates"
	listenPort    = "amyang.xyz:443"
	secretKey     = "CuteSweetie"
	loggerSkipper = func(ctx *sgo.Context) bool {
		if len(ctx.Path()) > 8 && ctx.Path()[0:8] == "/static/" {
			return true
		}
		return false
	}
	jwtSkipper = func(ctx *sgo.Context) bool {
		if ctx.Path() == "/api" && ctx.Method() == "POST" {
			return false
		}
		return true
	}
)

func main() {
	app := sgo.New()
	app.SetTemplates(tplDir, template.FuncMap{})
	app.USE(middlewares.Logger(os.Stdout, loggerSkipper))
	app.USE(middlewares.JWT("Header", secretKey, jwtSkipper))

	app.Any("/", home)
	app.GET("/static/*files", static)
	app.GET("/api", biu)
	app.POST("/api", hello)
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

func hello(ctx *sgo.Context) error {
	usr := ctx.Get("userInfo").(*jwt.Token)
	claims := usr.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return ctx.Text(200, "Hello "+name)
}

func usr(ctx *sgo.Context) error {
	return ctx.Text(200, "Welcome home, "+ctx.Param("user"))
}
