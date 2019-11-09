package middlewares

import "github.com/AmyangXYZ/sweetygo"

// Skipper defines your own function to skip a middleware.
// Returning true to skip.
type Skipper func(*sweetygo.Context) bool

// DefaultSkipper returns false to execute this middleware for all pages.
func DefaultSkipper(*sweetygo.Context) bool {
	return false
}
