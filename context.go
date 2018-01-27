package sweetygo

import (
	"net/http"
	"net/url"
)

// Context provide a HTTP context for SweetyGo.
type Context struct {
	sg       *SweetyGo
	Req      *http.Request
	Resp     *responseWriter
	handlers []HandlerFunc
}

// NewContext .
func NewContext(s *SweetyGo) *Context {
	c := &Context{}
	c.sg = s
	c.handlers = make([]HandlerFunc, len(s.middlewares)+1)
	copy(c.handlers, s.middlewares)
	return c
}

// Init the context gotten from sync pool.
func (c *Context) Init(w http.ResponseWriter, r *http.Request) {
	c.Resp = &responseWriter{w, 0}
	c.Req = r
	c.handlers = c.handlers[:len(c.sg.middlewares)]
}

// Next execute next middleware or router.
func (c *Context) Next() {
	n := len(c.handlers)
	switch {
	case n > 1:
		c.handlers[0](c)
		c.handlers = c.handlers[1:]
		c.Next()
	case n == 1:
		c.handlers[0](c)
	default:
		return
	}
}

// Params returns route params
func (c *Context) Params() url.Values {
	c.Req.ParseForm()
	return c.Req.Form
}
