package sgo

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

// HandlerFunc context handler func
type HandlerFunc func(*Context) error

// PreflightHandler is a dummy handler that handles preflight request when CORS
var PreflightHandler = func(ctx *Context) error { return ctx.Text(200, "") }

// SGo is Suuuuuuuuper Sweetie!
type SGo struct {
	// Router is based on a radix/trie tree.
	Tree *Trie

	// Pool is for context.
	Pool *sync.Pool

	NotFoundHandler         HandlerFunc
	MethodNotAllowedHandler HandlerFunc

	// Built-in templates render.
	Templates *Templates

	Middlewares []HandlerFunc
}

// New SGo App.
func New() *SGo {
	tree := &Trie{
		component: "/",
		methods:   make(map[string]HandlerFunc),
	}
	sg := &SGo{Tree: tree,
		NotFoundHandler:         NotFoundHandler,
		MethodNotAllowedHandler: MethodNotAllowedHandler,
		Middlewares:             make([]HandlerFunc, 0),
	}
	sg.Pool = &sync.Pool{
		New: func() interface{} {
			return NewContext(nil, nil, sg)
		},
	}
	return sg
}

// USE middlewares for SGo
func (sg *SGo) USE(middlewares ...HandlerFunc) {
	for i := range middlewares {
		if middlewares[i] != nil {
			sg.Middlewares = append(sg.Middlewares, middlewares[i])
		}
	}
}

// SetTemplates set the templates dir and funcMap.
func (sg *SGo) SetTemplates(tplDir string, funcMap template.FuncMap) {
	sg.Templates = NewTemplates(tplDir, funcMap)
}

// Run at the given addr
func (sg *SGo) Run(addr string) error {
	// logo, _ := base64.StdEncoding.DecodeString("XOKUgi/ilZTilZDilZfilKwg4pSs4pSM4pSA4pSQ4pSM4pSA4pSQ4pSM4pSs4pSQ4pSsIOKUrOKVlOKVkOKVl+KUjOKUgOKUkFzilIIvCuKUgCDilIDilZrilZDilZfilILilILilILilJzilKQg4pSc4pSkICDilIIg4pSU4pSs4pSY4pWRIOKVpuKUgiDilILilIAg4pSACi/ilIJc4pWa4pWQ4pWd4pSU4pS04pSY4pSU4pSA4pSY4pSU4pSA4pSYIOKUtCAg4pS0IOKVmuKVkOKVneKUlOKUgOKUmC/ilIJcCg==")
	// fmt.Println(string(logo))
	fmt.Printf("*SGo* -- Listen on %s\n", addr)
	return http.ListenAndServe(addr, sg)
}

// RunOverTLS .
func (sg *SGo) RunOverTLS(addr, certFile, keyFile string) error {
	// logo, _ := base64.StdEncoding.DecodeString("XOKUgi/ilZTilZDilZfilKwg4pSs4pSM4pSA4pSQ4pSM4pSA4pSQ4pSM4pSs4pSQ4pSsIOKUrOKVlOKVkOKVl+KUjOKUgOKUkFzilIIvCuKUgCDilIDilZrilZDilZfilILilILilILilJzilKQg4pSc4pSkICDilIIg4pSU4pSs4pSY4pWRIOKVpuKUgiDilILilIAg4pSACi/ilIJc4pWa4pWQ4pWd4pSU4pS04pSY4pSU4pSA4pSY4pSU4pSA4pSYIOKUtCAg4pS0IOKVmuKVkOKVneKUlOKUgOKUmC/ilIJcCg==")
	// fmt.Println(string(logo))
	fmt.Printf("*SGo* -- Listen on %s (TLS)\n", addr)
	return http.ListenAndServeTLS(addr, certFile, keyFile, sg)
}
