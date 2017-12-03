package sweetygo

import (
	"fmt"
	"net/http"
	"strings"
)

// Router is based on Radix Tree
// rootHandler is for runMiddleware
type Router struct {
	tree        *node
	rootHandler HandlerFunc
	middlewares []HandlerFunc
}

// New Router
func New(root HandlerFunc) *Router {
	node := node{
		children:  make([]*node, 0),
		prefix:    "/",
		hasParams: false,
		methods:   make(map[string]HandlerFunc),
	}
	return &Router{tree: &node,
		middlewares: make([]HandlerFunc, 0),
		rootHandler: root}
}

// USE middlewares for router
func (r *Router) USE(middleware ...HandlerFunc) {
	r.middlewares = append(r.middlewares, middleware...)
}

func runMiddleware(w http.ResponseWriter, req *http.Request, middleware Middleware) {
	middleware.ServeHTTP(w, req)
}

// build a middleware list
func mwareList(middleware []HandlerFunc, handler HandlerFunc) Middleware {
	var next Middleware

	if len(middleware) == 0 {
		return finalHandler(handler)
	} else if len(middleware) > 1 {
		next = mwareList(middleware[1:], handler)
	} else {
		next = finalHandler(handler)
	}

	return Middleware{middleware[0], &next}
}

func finalHandler(handler HandlerFunc) Middleware {
	return Middleware{
		HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			handler(w, r, next)
		}),
		&Middleware{}}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form

	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if handler := node.methods[req.Method]; handler != nil {
		runMiddleware(w, req, mwareList(r.middlewares, handler))
	} else {
		runMiddleware(w, req, mwareList(r.middlewares, r.rootHandler))
	}
}

// RunServer at the given addr
func (r *Router) RunServer(addr string) {
	fmt.Printf("*SweetyGo* -- Listen on %s\n", addr)
	http.ListenAndServe(addr, r)
}

// Handle register custom METHOD request HandlerFunc
func (r *Router) Handle(method, path string, handler HandlerFunc) {
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/sweety/go'")
	}
	r.tree.addNode(method, path, handler)
}

// GET register GET request handler
func (r *Router) GET(path string, handler HandlerFunc) {
	r.Handle("GET", path, handler)
}

// HEAD register HEAD request handler
func (r *Router) HEAD(path string, handler HandlerFunc) {
	r.Handle("HEAD", path, handler)
}

// OPTIONS register OPTIONS request handler
func (r *Router) OPTIONS(path string, handler HandlerFunc) {
	r.Handle("OPTIONS", path, handler)
}

// POST register POST request handler
func (r *Router) POST(path string, handler HandlerFunc) {
	r.Handle("POST", path, handler)
}

// PUT register PUT request handler
func (r *Router) PUT(path string, handler HandlerFunc) {
	r.Handle("PUT", path, handler)
}

// PATCH register PATCH request HandlerFunc
func (r *Router) PATCH(path string, handler HandlerFunc) {
	r.Handle("PATCH", path, handler)
}

// DELETE register DELETE request handler
func (r *Router) DELETE(path string, handler HandlerFunc) {
	r.Handle("DELETE", path, handler)
}
