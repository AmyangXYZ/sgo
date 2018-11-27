package sweetygo

import (
	"net/http"
	"strings"
)

// HTTP methods
const (
	CONNECT = http.MethodConnect
	DELETE  = http.MethodDelete
	GET     = http.MethodGet
	HEAD    = http.MethodHead
	OPTIONS = http.MethodOptions
	PATCH   = http.MethodPatch
	POST    = http.MethodPost
	PUT     = http.MethodPut
	TRACE   = http.MethodTrace
)

var methods = [...]string{
	http.MethodConnect,
	http.MethodDelete,
	http.MethodGet,
	http.MethodHead,
	http.MethodOptions,
	http.MethodPatch,
	http.MethodPost,
	http.MethodPut,
	http.MethodTrace,
}

// NotFoundHandler .
func NotFoundHandler(ctx *Context) {
	http.NotFound(ctx.Resp, ctx.Req)
}

// MethodNotAllowedHandler .
func MethodNotAllowedHandler(ctx *Context) {
	http.Error(ctx.Resp, "Method Not Allowed", 405)
}

func (sg *SweetyGo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := sg.Pool.Get().(*Context)
	ctx.Init(w, r)
	ctx.Req.ParseForm()
	node := sg.Tree.Search(strings.Split(r.URL.Path, "/")[1:], ctx.Params())
	if node != nil && node.methods[r.Method] != nil {
		ctx.handlers = append(ctx.handlers, node.methods[r.Method])
	} else if node != nil && node.methods[r.Method] == nil {
		ctx.handlers = append(ctx.handlers, sg.MethodNotAllowedHandler)
	} else {
		ctx.handlers = append(ctx.handlers, sg.NotFoundHandler)
	}

	ctx.Next()
	sg.Pool.Put(ctx)
}

// Handle register custom METHOD request HandlerFunc
func (sg *SweetyGo) Handle(method, path string, handler HandlerFunc) {
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/sweety/go'")
	}
	sg.Tree.Insert(method, path, handler)
}

// GET register GET request handler
func (sg *SweetyGo) GET(path string, handler HandlerFunc) {
	sg.Handle(GET, path, handler)
}

// HEAD register HEAD request handler
func (sg *SweetyGo) HEAD(path string, handler HandlerFunc) {
	sg.Handle(HEAD, path, handler)
}

// OPTIONS register OPTIONS request handler
func (sg *SweetyGo) OPTIONS(path string, handler HandlerFunc) {
	sg.Handle(OPTIONS, path, handler)
}

// POST register POST request handler
func (sg *SweetyGo) POST(path string, handler HandlerFunc) {
	sg.Handle(POST, path, handler)
}

// PUT register PUT request handler
func (sg *SweetyGo) PUT(path string, handler HandlerFunc) {
	sg.Handle(PUT, path, handler)
}

// PATCH register PATCH request HandlerFunc
func (sg *SweetyGo) PATCH(path string, handler HandlerFunc) {
	sg.Handle(PATCH, path, handler)
}

// DELETE register DELETE request handler
func (sg *SweetyGo) DELETE(path string, handler HandlerFunc) {
	sg.Handle(DELETE, path, handler)
}

// CONNECT register CONNECT request handler
func (sg *SweetyGo) CONNECT(path string, handler HandlerFunc) {
	sg.Handle(CONNECT, path, handler)
}

// TRACE register TRACE request handler
func (sg *SweetyGo) TRACE(path string, handler HandlerFunc) {
	sg.Handle(TRACE, path, handler)
}

// Any register any method handler
func (sg *SweetyGo) Any(path string, handler HandlerFunc) {
	for _, m := range methods {
		sg.Handle(m, path, handler)
	}
}
