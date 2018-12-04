package middlewares

import (
	"bufio"
	"compress/gzip"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/AmyangXYZ/sweetygo"
)

// Gzip compress.
func Gzip(level int, skipper Skipper) sweetygo.HandlerFunc {
	return func(ctx *sweetygo.Context) error {
		if skipper(ctx) == true {
			ctx.Next()
			return nil
		}
		if !strings.Contains(ctx.Req.Header.Get("Accept-Encoding"), "gzip") {
			ctx.Next()
			return nil
		}
		ctx.Resp.Header().Set("Content-Encoding", "gzip")
		w := ctx.Resp.ResponseWriter
		gw, _ := gzip.NewWriterLevel(w, level)
		defer gw.Close()
		ctx.Resp.ResponseWriter = &gzipResponseWriter{Writer: gw, ResponseWriter: w}
		ctx.Next()
		return nil
	}
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// func (w *gzipResponseWriter) WriteHeader(c int) {
// 	w.ResponseWriter.Header().Del("Content-Length")
// 	w.ResponseWriter.WriteHeader(c)
// }

func (w *gzipResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
