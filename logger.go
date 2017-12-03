package sweetygo

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Logger is a built-in middleware
func Logger() HandlerFunc {
	logger := log.New(os.Stdout, "*SweetyGo*", 0)
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		start := time.Now()
		rw := &responseWriteReader{w, 0}
		next(rw, r)
		end := time.Since(start)
		logger.Printf(" -- %s - [%v] \"%s %s\" %d - %v",
			strings.Split(r.RemoteAddr, ":")[0], start.Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, rw.status, end)
	}
}
