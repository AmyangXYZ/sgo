package sweetygo

import (
	"fmt"
	"net/http"
	"sync"
)

// HandlerFunc context handler func
type HandlerFunc func(*Context)

// Middleware handler
type Middleware interface{}

// SweetyGo is Suuuuuuuuper Sweetie!
type SweetyGo struct {
	tree            *Trie
	pool            sync.Pool
	notFoundHandler HandlerFunc
	middlewares     []HandlerFunc
}

// New SweetyGo App
func New() *SweetyGo {
	tree := &Trie{
		component: "/",
		methods:   make(map[string]HandlerFunc),
	}
	s := &SweetyGo{tree: tree,
		notFoundHandler: NotFoundHandler,
		middlewares:     make([]HandlerFunc, 0),
	}
	s.pool = sync.Pool{
		New: func() interface{} {
			return NewContext(nil, nil, s)
		},
	}
	return s
}

// USE middlewares for SweetyGo
func (s *SweetyGo) USE(m ...Middleware) {
	for i := range m {
		if m[i] != nil {
			s.middlewares = append(s.middlewares, wrapMiddleware(m[i]))
		}
	}
}

// RunServer at the given addr
func (s *SweetyGo) RunServer(addr string) {
	fmt.Printf("*SweetyGo* -- Listen on %s\n", addr)
	http.ListenAndServe(addr, s)
}

// wrapMiddleware wraps middleware.
func wrapMiddleware(m Middleware) HandlerFunc {
	switch m := m.(type) {
	case HandlerFunc:
		return m
	case func(*Context):
		return m
	case http.Handler, http.HandlerFunc:
		return WrapHandlerFunc(func(c *Context) {
			m.(http.Handler).ServeHTTP(c.Resp, c.Req)
		})
	case func(http.ResponseWriter, *http.Request):
		return WrapHandlerFunc(func(c *Context) {
			m(c.Resp, c.Req)
		})
	default:
		panic("unknown middleware")
	}
}

// WrapHandlerFunc wrap for context handler chain
func WrapHandlerFunc(h HandlerFunc) HandlerFunc {
	return func(c *Context) {
		h(c)
		c.Next()
	}
}
