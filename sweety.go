package sweetygo

import "net/http"

type Handler interface {
	ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

type Middleware struct {
	handler Handler
	next    *Middleware
}

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h(w, r, next)
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.handler.ServeHTTP(w, r, m.next.ServeHTTP)
}
