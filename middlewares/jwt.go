package middlewares

import (
	"github.com/AmyangXYZ/sweetygo"
	"github.com/dgrijalva/jwt-go"
)

// DefaultJWTConfig is a built-in middleware with default
// configuration for authentication.
type DefaultJWTConfig struct {
	// Skipper defines a function to skip middleware.
	Skipper Skipper

	// Token name stored in cookie.
	// Default is `sgtoken`.
	Name string

	// Default is HS256.
	SigningMethod jwt.SigningMethod

	// Secret key.
	Keyfunc jwt.Keyfunc

	// store in cookie or header.
	// Cookie is safer, https://stormpath.com/blog/where-to-store-your-jwts-cookies-vs-html5-web-storage.
	Store string

	// Context key to store user information from the token into context.
	// Default is usrToken.
	ContextKey string

	// Claims are extendable claims data defining token content.
	// Default is jwt.MapClaims.
	Claims jwt.Claims
}

// JWT implements sweetygo.HandlerFunc.
func JWT(store, key string, skipper Skipper) sweetygo.HandlerFunc {
	J := &DefaultJWTConfig{
		Skipper:       skipper,
		Name:          "SG_Token",
		SigningMethod: jwt.SigningMethodHS256,
		Keyfunc: func(t *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
		Store:      store,
		ContextKey: "userInfo",
		Claims:     jwt.MapClaims{},
	}
	return func(ctx *sweetygo.Context) {
		if J.Skipper(ctx) == true {
			ctx.Next()
			return
		}
		auth := J.extractorJWT(ctx)
		token, err := jwt.Parse(auth, J.Keyfunc)
		if err == nil && token.Valid {
			ctx.Set(J.ContextKey, token)
			ctx.Next()
			return
		}
		ctx.Error("Unauthorized access to this resource", 401)
	}
}

func (J *DefaultJWTConfig) extractorJWT(ctx *sweetygo.Context) string {
	switch J.Store {
	case "Cookie":
		return ctx.GetCookie(J.Name)
	case "Header":
		auth := ctx.Req.Header.Get("Authorization")
		if len(auth) > len(J.Name)+1 {
			return auth[len(J.Name)+1:]
		}
		return ""
	default:
		return ""
	}
}
