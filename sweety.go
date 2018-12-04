package sweetygo

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"github.com/lucas-clemente/quic-go/h2quic"
)

// HandlerFunc context handler func
type HandlerFunc func(*Context) error

// SweetyGo is Suuuuuuuuper Sweetie!
type SweetyGo struct {
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

// New SweetyGo App.
func New() *SweetyGo {
	tree := &Trie{
		component: "/",
		methods:   make(map[string]HandlerFunc),
	}
	sg := &SweetyGo{Tree: tree,
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

// USE middlewares for SweetyGo
func (sg *SweetyGo) USE(middlewares ...HandlerFunc) {
	for i := range middlewares {
		if middlewares[i] != nil {
			sg.Middlewares = append(sg.Middlewares, middlewares[i])
		}
	}
}

// SetTemplates set the templates dir and funcMap.
func (sg *SweetyGo) SetTemplates(tplDir string, funcMap template.FuncMap) {
	sg.Templates = NewTemplates(tplDir, funcMap)
}

// Run at the given addr
func (sg *SweetyGo) Run(addr string) error {
	logo, _ := base64.StdEncoding.DecodeString("XOKUgi/ilZTilZDilZfilKwg4pSs4pSM4pSA4pSQ4pSM4pSA4pSQ4pSM4pSs4pSQ4pSsIOKUrOKVlOKVkOKVl+KUjOKUgOKUkFzilIIvCuKUgCDilIDilZrilZDilZfilILilILilILilJzilKQg4pSc4pSkICDilIIg4pSU4pSs4pSY4pWRIOKVpuKUgiDilILilIAg4pSACi/ilIJc4pWa4pWQ4pWd4pSU4pS04pSY4pSU4pSA4pSY4pSU4pSA4pSYIOKUtCAg4pS0IOKVmuKVkOKVneKUlOKUgOKUmC/ilIJcCg==")
	fmt.Println(string(logo))
	fmt.Printf("*SweetyGo* -- Listen on %s\n", addr)
	return http.ListenAndServe(addr, sg)
}

// RunOverTLS .
func (sg *SweetyGo) RunOverTLS(addr, certFile, keyFile string) error {
	logo, _ := base64.StdEncoding.DecodeString("XOKUgi/ilZTilZDilZfilKwg4pSs4pSM4pSA4pSQ4pSM4pSA4pSQ4pSM4pSs4pSQ4pSsIOKUrOKVlOKVkOKVl+KUjOKUgOKUkFzilIIvCuKUgCDilIDilZrilZDilZfilILilILilILilJzilKQg4pSc4pSkICDilIIg4pSU4pSs4pSY4pWRIOKVpuKUgiDilILilIAg4pSACi/ilIJc4pWa4pWQ4pWd4pSU4pS04pSY4pSU4pSA4pSY4pSU4pSA4pSYIOKUtCAg4pS0IOKVmuKVkOKVneKUlOKUgOKUmC/ilIJcCg==")
	fmt.Println(string(logo))
	fmt.Printf("*SweetyGo* -- Listen on %s (TLS)\n", addr)
	return http.ListenAndServeTLS(addr, certFile, keyFile, sg)
}

// RunOverQUIC .
func (sg *SweetyGo) RunOverQUIC(addr, certFile, keyFile string) error {
	logo, _ := base64.StdEncoding.DecodeString("XOKUgi/ilZTilZDilZfilKwg4pSs4pSM4pSA4pSQ4pSM4pSA4pSQ4pSM4pSs4pSQ4pSsIOKUrOKVlOKVkOKVl+KUjOKUgOKUkFzilIIvCuKUgCDilIDilZrilZDilZfilILilILilILilJzilKQg4pSc4pSkICDilIIg4pSU4pSs4pSY4pWRIOKVpuKUgiDilILilIAg4pSACi/ilIJc4pWa4pWQ4pWd4pSU4pS04pSY4pSU4pSA4pSY4pSU4pSA4pSYIOKUtCAg4pS0IOKVmuKVkOKVneKUlOKUgOKUmC/ilIJcCg==")
	fmt.Println(string(logo))
	fmt.Printf("*SweetyGo* -- Listen on %s (HTTP2+QUIC)\n", addr)
	return h2quic.ListenAndServe(addr, certFile, keyFile, sg)
}
