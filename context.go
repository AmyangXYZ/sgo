package sweetygo

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

// Context provide a HTTP context for SweetyGo.
type Context struct {
	sg           *SweetyGo
	Req          *http.Request
	Resp         *responseWriter
	handlers     []HandlerFunc
	handlerState int
}

// NewContext .
func NewContext(w http.ResponseWriter, r *http.Request, sg *SweetyGo) *Context {
	ctx := &Context{}
	ctx.sg = sg
	ctx.handlers = make([]HandlerFunc, len(sg.Middlewares), len(sg.Middlewares)+3)
	copy(ctx.handlers, sg.Middlewares)
	ctx.Init(w, r)
	return ctx
}

// Init the context gotten from sync pool.
func (ctx *Context) Init(w http.ResponseWriter, r *http.Request) {
	ctx.Resp = &responseWriter{w, 0}
	ctx.Req = r
	ctx.handlers = ctx.handlers[:len(ctx.sg.Middlewares)]
	ctx.handlerState = 0
}

// Next execute next middleware or router.
func (ctx *Context) Next() {
	if ctx.handlerState < len(ctx.handlers) {
		i := ctx.handlerState
		ctx.handlerState++
		ctx.handlers[i](ctx)
	}
}

// Params returns route params
func (ctx *Context) Params() url.Values {
	ctx.Req.ParseForm()
	return ctx.Req.Form
}

// Method .
func (ctx *Context) Method() string {
	return ctx.Req.Method
}

// FormFile gets file from request.
func (ctx *Context) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return ctx.Req.FormFile(key)
}

// SaveFile saves the form file.
func (ctx *Context) SaveFile(name, savePath string) error {
	fr, _, err := ctx.FormFile(name)
	if err != nil {
		return err
	}
	defer fr.Close()

	fw, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer fw.Close()

	_, err = io.Copy(fw, fr)
	return err
}

// Error .
func (ctx *Context) Error(error string, code int) {
	http.Error(ctx.Resp, error, code)
}

// Write Response.
func (ctx *Context) Write(data []byte) (n int, err error) {
	return ctx.Resp.Write(data)
}

// HTML response HTML data.
func (ctx *Context) HTML(code int, body string) {
	ctx.Resp.Header().Set("Content-Type", "text/html")
	ctx.Resp.WriteHeader(code)
	ctx.Resp.Write([]byte(body))
}

// JSON response JSON data.
func (ctx *Context) JSON(code int, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		ctx.Resp.WriteHeader(http.StatusInternalServerError)
		ctx.Error("JSON Marshal Error", 500)
	}
	ctx.Resp.Header().Set("Content-Type", "application/json")
	ctx.Resp.WriteHeader(code)
	ctx.Resp.Write(data)
}

// Redirect redirects the request
func (ctx *Context) Redirect(url string, code int) {
	http.Redirect(ctx.Resp, ctx.Req, url, code)
}

// URL returns URL string.
func (ctx *Context) URL() string {
	return ctx.Req.URL.Path
}

// Referer returns request referer.
func (ctx *Context) Referer() string {
	return ctx.Req.Header.Get("Referer")
}

// UserAgent returns http request UserAgent
func (ctx *Context) UserAgent() string {
	return ctx.Req.Header.Get("User-Agent")
}
