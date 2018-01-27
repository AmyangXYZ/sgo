package sweetygo

import (
	"net/http"
	"net/url"
)

// Context provide a HTTP context for SweetyGo.
type Context struct {
	sg           *SweetyGo
	Req          *http.Request
	Resp         *responseWriter
	handlers     []HandlerFunc
	handlerState int
}

// NewContext .
func NewContext(w http.ResponseWriter, r *http.Request, s *SweetyGo) *Context {
	c := &Context{}
	c.sg = s
	c.handlers = make([]HandlerFunc, len(s.middlewares), len(s.middlewares)+3)
	copy(c.handlers, s.middlewares)
	c.Init(w, r)
	return c
}

// Init the context gotten from sync pool.
func (c *Context) Init(w http.ResponseWriter, r *http.Request) {
	c.Resp = &responseWriter{w, 0}
	c.Req = r
	c.handlers = c.handlers[:len(c.sg.middlewares)]
	c.handlerState = 0
}

// Next execute next middleware or router.
func (c *Context) Next() {
	if c.handlerState < len(c.handlers) {
		i := c.handlerState
		c.handlerState++
		c.handlers[i](c)
	}
}

// Params returns route params
func (c *Context) Params() url.Values {
	c.Req.ParseForm()
	return c.Req.Form
}
