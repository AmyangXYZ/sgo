package middlewares

import "github.com/AmyangXYZ/sgo"

// Skipper defines your own function to skip a middleware.
// Returning true to skip.
type Skipper func(*sgo.Context) bool

// DefaultSkipper returns false to execute this middleware for all pages.
func DefaultSkipper(*sgo.Context) bool {
	return false
}
