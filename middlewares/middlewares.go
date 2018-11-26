package middlewares

import "github.com/AmyangXYZ/sweetygo"

// Skipper defines a function to skip middleware. Returning true skips processing
// the middleware.
type Skipper func(*sweetygo.Context) bool

// DefaultSkipper returns false which processes the middleware.
func DefaultSkipper(*sweetygo.Context) bool {
	return false
}
