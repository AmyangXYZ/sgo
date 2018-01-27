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
			return NewContext(s)
		},
	}
	return s
}

// USE middlewares for SweetyGo
func (s *SweetyGo) USE(middleware ...HandlerFunc) {
	s.middlewares = append(s.middlewares, middleware...)
}

// RunServer at the given addr
func (s *SweetyGo) RunServer(addr string) {
	fmt.Printf("*SweetyGo* -- Listen on %s\n", addr)
	http.ListenAndServe(addr, s)
}
