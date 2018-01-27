package sweetygo

import (
	"net/http"
	"strings"
)

// NotFoundHandler .
func NotFoundHandler(c *Context) {
	http.NotFound(c.Resp, c.Req)
}

func (s *SweetyGo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := s.pool.Get().(*Context)
	c.Init(w, r)

	node := s.tree.Search(strings.Split(r.URL.Path, "/")[1:], c.Params())
	if node != nil && node.methods[r.Method] != nil {
		c.handlers = append(c.handlers, node.methods[r.Method])
	} else {
		c.handlers = append(c.handlers, s.notFoundHandler)
	}

	c.Next()
	s.pool.Put(c)
}

// Handle register custom METHOD request HandlerFunc
func (s *SweetyGo) Handle(method, path string, handler HandlerFunc) {
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/sweety/go'")
	}
	s.tree.Insert(method, path, handler)
}

// GET register GET request handler
func (s *SweetyGo) GET(path string, handler HandlerFunc) {
	s.Handle("GET", path, handler)
}

// HEAD register HEAD request handler
func (s *SweetyGo) HEAD(path string, handler HandlerFunc) {
	s.Handle("HEAD", path, handler)
}

// OPTIONS register OPTIONS request handler
func (s *SweetyGo) OPTIONS(path string, handler HandlerFunc) {
	s.Handle("OPTIONS", path, handler)
}

// POST register POST request handler
func (s *SweetyGo) POST(path string, handler HandlerFunc) {
	s.Handle("POST", path, handler)
}

// PUT register PUT request handler
func (s *SweetyGo) PUT(path string, handler HandlerFunc) {
	s.Handle("PUT", path, handler)
}

// PATCH register PATCH request HandlerFunc
func (s *SweetyGo) PATCH(path string, handler HandlerFunc) {
	s.Handle("PATCH", path, handler)
}

// DELETE register DELETE request handler
func (s *SweetyGo) DELETE(path string, handler HandlerFunc) {
	s.Handle("DELETE", path, handler)
}
