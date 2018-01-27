package sweetygo

import (
	"fmt"
	"net/http"
	"sync"
)

// HandlerFunc context handler func
type HandlerFunc func(*Context)

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
	sg := &SweetyGo{tree: tree,
		notFoundHandler: NotFoundHandler,
		middlewares:     make([]HandlerFunc, 0),
	}
	sg.pool = sync.Pool{
		New: func() interface{} {
			return NewContext(nil, nil, sg)
		},
	}
	return sg
}

// USE middlewares for SweetyGo
func (sg *SweetyGo) USE(middlewares ...HandlerFunc) {
	for i := range middlewares {
		if middlewares[i] != nil {
			sg.middlewares = append(sg.middlewares, middlewares[i])
		}
	}
}

// RunServer at the given addr
func (sg *SweetyGo) RunServer(addr string) {
	fmt.Printf("*SweetyGo* -- Listen on %s\n", addr)
	http.ListenAndServe(addr, sg)
}
