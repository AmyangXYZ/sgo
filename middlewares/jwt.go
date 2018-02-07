package middlewares

import (
	"strings"

	"github.com/AmyangXYZ/sweetygo"
	"github.com/dgrijalva/jwt-go"
)

// JWTConfig is a built-in middleware with default
// configuration for authentication.
type JWTConfig struct {
	// Token name stored in cookie.
	// Default is `sgtoken`.
	Name string

	// Default is HS256.
	SigningMethod jwt.SigningMethod

	// Secret key.
	Keyfunc jwt.Keyfunc

	// Path and method who require JWT middleware.
	// Example:
	// requiredJWTMap = map[string]string{
	// 	"/api/1":    "POST",
	// 	"/api/2":    "ALL",
	// 	"/api/3":    "!GET",
	// 	"/secret/*": "ALL",
	// }
	RequiredMap map[string]string

	// Context key to store user information from the token into context.
	// Default is user.
	ContextKey string

	// Claims are extendable claims data defining token content.
	// Default is jwt.MapClaims.
	Claims jwt.Claims
}

// JWT implements sweetygo.HandlerFunc.
func JWT(key string, requiredmap map[string]string) sweetygo.HandlerFunc {
	J := &JWTConfig{
		Name:          "sgtoken",
		SigningMethod: jwt.SigningMethodHS256,
		Keyfunc: func(t *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
		RequiredMap: requiredmap,
		ContextKey:  "user",
		Claims:      jwt.MapClaims{},
	}
	return func(ctx *sweetygo.Context) {
		for path, method := range J.RequiredMap {
			if (path[len(path)-1] == '*' && strings.HasPrefix(ctx.URL(), path[0:len(path)-1])) ||
				ctx.URL() == path {
				if ctx.Method() == method || method == "ALL" ||
					(method[0] == '!' && ctx.Method() != method[1:]) {
					auth := ctx.GetCookie(J.Name)
					token, err := jwt.Parse(auth, J.Keyfunc)
					if err == nil && token.Valid {
						ctx.Set(J.ContextKey, token)
						ctx.Next()
						return
					}
					ctx.Error("Unauthorized access to this resource", 401)
				}
			}
		}
		ctx.Next()
		return
	}
}
