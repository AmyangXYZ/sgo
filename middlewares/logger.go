package middlewares

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/AmyangXYZ/sweetygo"
)

// Logger implements SweetyGo.Middleware's interface.
func Logger() sweetygo.HandlerFunc {
	logger := log.New(os.Stdout, "*SweetyGo*", 0)
	return func(ctx *sweetygo.Context) {
		start := time.Now()
		ctx.Next()
		end := time.Since(start)
		logger.Printf(" -- %s - [%v] \"%s %s %d\" \"%s\" \"%s\" - %v",
			strings.Split(ctx.Req.RemoteAddr, ":")[0], start.Format("2006-01-02 15:04:05"),
			ctx.Method(), ctx.URL(), ctx.Status(), ctx.Referer(), ctx.UserAgent(), end)
	}
}
