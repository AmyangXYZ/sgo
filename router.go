package sweetygo

import (
	"net/http"
	"strings"
)

// NotFoundHandler .
func NotFoundHandler(ctx *Context) {
	http.NotFound(ctx.Resp, ctx.Req)
}

// MethodNotAllowedHandler .
func MethodNotAllowedHandler(ctx *Context) {
	http.Error(ctx.Resp, "Method Not Allowed", 405)
}

// Static .
func (sg *SweetyGo) Static(path, dir string) {
	StaticServer := func(ctx *Context) {
		staticHandle := http.StripPrefix(path,
			http.FileServer(http.Dir(dir)))
		staticHandle.ServeHTTP(ctx.Resp, ctx.Req)
	}
	sg.GET(path+"/*files", StaticServer)
}

func (sg *SweetyGo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := sg.Pool.Get().(*Context)
	ctx.Init(w, r)

	node := sg.Tree.Search(strings.Split(r.URL.Path, "/")[1:], ctx.ParseForm())
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
	sg.Handle("GET", path, handler)
}

// HEAD register HEAD request handler
func (sg *SweetyGo) HEAD(path string, handler HandlerFunc) {
	sg.Handle("HEAD", path, handler)
}

// OPTIONS register OPTIONS request handler
func (sg *SweetyGo) OPTIONS(path string, handler HandlerFunc) {
	sg.Handle("OPTIONS", path, handler)
}

// POST register POST request handler
func (sg *SweetyGo) POST(path string, handler HandlerFunc) {
	sg.Handle("POST", path, handler)
}

// PUT register PUT request handler
func (sg *SweetyGo) PUT(path string, handler HandlerFunc) {
	sg.Handle("PUT", path, handler)
}

// PATCH register PATCH request HandlerFunc
func (sg *SweetyGo) PATCH(path string, handler HandlerFunc) {
	sg.Handle("PATCH", path, handler)
}

// DELETE register DELETE request handler
func (sg *SweetyGo) DELETE(path string, handler HandlerFunc) {
	sg.Handle("DELETE", path, handler)
}
