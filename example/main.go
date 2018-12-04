package main

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/AmyangXYZ/sweetygo"
	"github.com/AmyangXYZ/sweetygo/middlewares"
	"github.com/dgrijalva/jwt-go"
)

var (
	tplDir        = "templates"
	listenPort    = ":16311"
	secretKey     = "CuteSweetie"
	loggerSkipper = func(ctx *sweetygo.Context) bool {
		if len(ctx.Path()) > 8 && ctx.Path()[0:8] == "/static/" {
			return true
		}
		return false
	}
	jwtSkipper = func(ctx *sweetygo.Context) bool {
		if ctx.Path() == "/api" && ctx.Method() == "POST" {
			return false
		}
		return true
	}
)

func main() {
	app := sweetygo.New()
	app.SetTemplates(tplDir, template.FuncMap{})
	app.USE(middlewares.Logger(os.Stdout, loggerSkipper))
	app.USE(middlewares.JWT("Header", secretKey, jwtSkipper))

	app.Any("/", home)
	app.GET("/static/*files", static)
	app.GET("/api", biu)
	app.POST("/api", hello)
	app.POST("/login", login)
	app.GET("/usr/:user", usr)

	app.Run(listenPort)

	// Or use QUIC
	// app.RunOverQUIC(listenPort, "fullchain.pem", "privkey.pem")
}

func home(ctx *sweetygo.Context) error {
	ctx.Set("baby", "Sweetie")
	return ctx.Render(200, "index")
}

func static(ctx *sweetygo.Context) error {
	staticHandle := http.StripPrefix("/static",
		http.FileServer(http.Dir("./static")))
	staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
	return nil
}

func biu(ctx *sweetygo.Context) error {
	return ctx.Text(200, "biu")
}

func login(ctx *sweetygo.Context) error {
	usr := ctx.Param("usr")
	pwd := ctx.Param("pwd")
	if usr == "Amyang" && pwd == "biu" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Amyang"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return err
		}
		return ctx.JSON(200, 1, "success", map[string]string{"SG_Token": t})
	}
	return ctx.JSON(200, 0, "username or password error", nil)
}

func api(ctx *sweetygo.Context) error {
	return ctx.JSON(200, 1, "success", map[string]string{"version": "1.1"})
}

func hello(ctx *sweetygo.Context) error {
	usr := ctx.Get("userInfo").(*jwt.Token)
	claims := usr.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return ctx.Text(200, "Hello "+name)
}

func usr(ctx *sweetygo.Context) error {
	return ctx.Text(200, "Welcome home, "+ctx.Param("user"))
}
