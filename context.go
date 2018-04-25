package sweetygo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
)

// Context provide a HTTP context for SweetyGo.
type Context struct {
	sg           *SweetyGo
	Req          *http.Request
	Resp         *responseWriter
	handlers     []HandlerFunc
	store        map[string]interface{}
	storeMutex   *sync.RWMutex
	handlerState int
}

// NewContext .
func NewContext(w http.ResponseWriter, r *http.Request, sg *SweetyGo) *Context {
	ctx := &Context{}
	ctx.sg = sg
	ctx.storeMutex = new(sync.RWMutex)
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
	ctx.storeMutex.Lock()
	ctx.store = nil
	ctx.storeMutex.Unlock()
}

// Next execute next middleware or router.
func (ctx *Context) Next() {
	if ctx.handlerState < len(ctx.handlers) {
		i := ctx.handlerState
		ctx.handlerState++
		ctx.handlers[i](ctx)
	}
	return
}

// Set var in context.
func (ctx *Context) Set(key string, val interface{}) {
	ctx.storeMutex.Lock()
	if ctx.store == nil {
		ctx.store = make(map[string]interface{})
	}
	ctx.store[key] = val
	ctx.storeMutex.Unlock()
}

// Get data in context.
func (ctx *Context) Get(key string) interface{} {
	ctx.storeMutex.RLock()
	v := ctx.store[key]
	ctx.storeMutex.RUnlock()
	return v
}

// Gets all data in context.
func (ctx *Context) Gets() map[string]interface{} {
	ctx.storeMutex.RLock()
	vals := make(map[string]interface{})
	for k, v := range ctx.store {
		vals[k] = v
	}
	ctx.storeMutex.RUnlock()
	return vals
}

// SetCookie is used for jwt.
func (ctx *Context) SetCookie(name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   0,
	}
	ctx.Resp.Header().Add("Set-Cookie", cookie.String())
}

// GetCookie .
func (ctx *Context) GetCookie(name string) string {
	cookie, err := ctx.Req.Cookie(name)
	if err != nil {
		return ""
	}
	v, _ := url.QueryUnescape(cookie.Value)
	return v
}

// Params returns all params
func (ctx *Context) Params() url.Values {
	return ctx.Req.Form
}

// Param returns specific params
func (ctx *Context) Param(key string) string {
	if ctx.Params()[key] != nil {
		return ctx.Params()[key][0]
	}
	return ""
}

// Method .
func (ctx *Context) Method() string {
	return ctx.Req.Method
}

//Status Code.
func (ctx *Context) Status() int {
	return ctx.Resp.status
}

// FormFile gets file from request.
func (ctx *Context) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return ctx.Req.FormFile(key)
}

// SaveFile saves the form file and
// returns the filename.
func (ctx *Context) SaveFile(name, saveDir string) (string, error) {
	fr, handle, err := ctx.FormFile(name)
	if err != nil {
		return "", err
	}
	defer fr.Close()
	fw, err := os.OpenFile(path.Join(saveDir, handle.Filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return "", err
	}
	defer fw.Close()

	_, err = io.Copy(fw, fr)
	return handle.Filename, err
}

// Error .
func (ctx *Context) Error(error string, code int) {
	http.Error(ctx.Resp, error, code)
	ctx.handlerState = len(ctx.handlers) // break handlers chain.
}

// Write Response.
func (ctx *Context) Write(data []byte) (n int, err error) {
	return ctx.Resp.Write(data)
}

// Text response text data.
func (ctx *Context) Text(code int, body string) {
	ctx.Resp.Header().Set("Content-Type", "text/plain; charset=utf-8")
	ctx.Resp.WriteHeader(code)
	ctx.Resp.Write([]byte(body))
}

// JSON response JSON data.
// {flag: 1, msg: "success", data: ...}
func (ctx *Context) JSON(code, flag int, msg string, data interface{}) {
	m := map[string]interface{}{
		"msg":  msg,
		"data": data,
		"flag": flag,
	}

	j, _ := json.Marshal(m)
	ctx.Resp.Header().Set("Content-Type", "application/json")
	ctx.Resp.WriteHeader(code)
	ctx.Resp.Write(j)
}

// JSONP return JSONP data.
func (ctx *Context) JSONP(code int, callback string, data interface{}) {
	j, _ := json.Marshal(data)
	ctx.Resp.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	ctx.Resp.WriteHeader(code)
	ctx.Resp.Write([]byte(callback + "("))
	ctx.Resp.Write(j)
	ctx.Resp.Write([]byte(");"))
}

// Render sweetygo.templates with stored data.
func (ctx *Context) Render(code int, tplname string) {
	buf := new(bytes.Buffer)
	err := ctx.sg.Templates.Render(buf, tplname, ctx.Gets())
	if err != nil {
		fmt.Println(err)
		ctx.Error("Render Error", 500)
		return
	}
	ctx.Resp.Header().Set("Content-Type", "text/html")
	ctx.Resp.WriteHeader(code)
	ctx.Resp.Write(buf.Bytes())
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
