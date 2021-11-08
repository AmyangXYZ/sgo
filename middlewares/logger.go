package middlewares

import (
	"io"
	"log"
	"strings"
	"time"

	"github.com/AmyangXYZ/sgo"
)

// Logger implements SGo.Middleware's interface.
func Logger(out io.Writer, skipper Skipper) sgo.HandlerFunc {
	logger := log.New(out, "*SGo*", 0)
	return func(ctx *sgo.Context) error {
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
