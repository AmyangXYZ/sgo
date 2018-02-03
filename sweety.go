package sweetygo

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"path"
	"sync"
)

// HandlerFunc context handler func
type HandlerFunc func(*Context)

// SweetyGo is Suuuuuuuuper Sweetie!
type SweetyGo struct {
	Tree                    *Trie
	Pool                    sync.Pool
	NotFoundHandler         HandlerFunc
	MethodNotAllowedHandler HandlerFunc
	Templates               *Templates
	Middlewares             []HandlerFunc
}

// New SweetyGo App.
// Deafault static and templates dir is
// rootDir+"static" and rootDir + "templates"
func New(rootDir string) *SweetyGo {
	tree := &Trie{
		component: "/",
		methods:   make(map[string]HandlerFunc),
	}
	sg := &SweetyGo{Tree: tree,
		NotFoundHandler:         NotFoundHandler,
		MethodNotAllowedHandler: MethodNotAllowedHandler,
		Templates:               NewTemplates(path.Join(rootDir, "templates")),
		Middlewares:             make([]HandlerFunc, 0),
	}
	sg.Pool = sync.Pool{
		New: func() interface{} {
			return NewContext(nil, nil, sg)
		},
	}
	sg.Static("/static", path.Join(rootDir, "static"))
	return sg
}

// USE middlewares for SweetyGo
func (sg *SweetyGo) USE(middlewares ...HandlerFunc) {
	for i := range middlewares {
		if middlewares[i] != nil {
			sg.Middlewares = append(sg.Middlewares, middlewares[i])
		}
	}
}

// RunServer at the given addr
func (sg *SweetyGo) RunServer(addr string) {
	logo, _ := base64.StdEncoding.DecodeString("XOKUgi/ilZTilZDilZfilKwg4pSs4pSM4pSA4pSQ4pSM4pSA4pSQ4pSM4pSs4pSQ4pSsIOKUrOKVlOKVkOKVl+KUjOKUgOKUkFzilIIvCuKUgCDilIDilZrilZDilZfilILilILilILilJzilKQg4pSc4pSkICDilIIg4pSU4pSs4pSY4pWRIOKVpuKUgiDilILilIAg4pSACi/ilIJc4pWa4pWQ4pWd4pSU4pS04pSY4pSU4pSA4pSY4pSU4pSA4pSYIOKUtCAg4pS0IOKVmuKVkOKVneKUlOKUgOKUmC/ilIJcCg==")
	fmt.Println(string(logo))
	fmt.Printf("*SweetyGo* -- Listen on %s\n", addr)
	http.ListenAndServe(addr, sg)
}
