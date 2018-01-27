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
	return func(c *Context) {
		start := time.Now()
		c.Next()
		end := time.Since(start)
		logger.Printf(" -- %s - [%v] \"%s %s\" %d - %v",
			strings.Split(c.Req.RemoteAddr, ":")[0], start.Format("2006-01-02 15:04:05"),
			c.Req.Method, c.Req.URL.Path, c.Resp.status, end)
	}
}
