package sweetygo

import (
	"log"
	"os"
	"strings"
	"time"
)

// Logger is a built-in middleware
func Logger(c *Context) {
	logger := log.New(os.Stdout, "*SweetyGo*", 0)
	start := time.Now()
	c.Next()
	logger.Printf(" -- %s - [%v] \"%s %s\" %d - %v",
		strings.Split(c.Req.RemoteAddr, ":")[0], start.Format("2006-01-02 15:04:05"),
		c.Req.Method, c.Req.URL.Path, c.Resp.status, time.Since(start))
}
