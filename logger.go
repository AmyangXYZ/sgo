package sweetygo

import (
	"log"
	"os"
	"strings"
	"time"
)

// Logger is a built-in middleware
func Logger() HandlerFunc {
	logger := log.New(os.Stdout, "*SweetyGo*", 0)
	return func(ctx *Context) {
		start := time.Now()
		ctx.Next()
		end := time.Since(start)
		logger.Printf(" -- %s - [%v] \"%s %s %d\" \"%s\" \"%s\" - %v",
			strings.Split(ctx.Req.RemoteAddr, ":")[0], start.Format("2006-01-02 15:04:05"),
			ctx.Method(), ctx.URL(), ctx.Resp.status, ctx.Referer(), ctx.UserAgent(), end)
	}
}
