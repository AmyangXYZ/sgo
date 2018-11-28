package middlewares

import (
	"io"
	"log"
	"strings"
	"time"

	"github.com/AmyangXYZ/sweetygo"
)

// Logger implements SweetyGo.Middleware's interface.
func Logger(out io.Writer, skipper Skipper) sweetygo.HandlerFunc {
	logger := log.New(out, "*SweetyGo*", 0)
	return func(ctx *sweetygo.Context) error {
		if skipper(ctx) == true {
			ctx.Next()
			return nil
		}
		start := time.Now()
		ctx.Next()
		end := time.Since(start)
		logger.Printf(" -- %s - [%v] \"%s %s %d\" \"%s\" \"%s\" - %v",
			strings.Split(ctx.Req.RemoteAddr, ":")[0], start.Format("2006-01-02 15:04:05"),
			ctx.Method(), ctx.Path(), ctx.Status(), ctx.Referer(), ctx.UserAgent(), end)
		return nil
	}
}
