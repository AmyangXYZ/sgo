package middlewares

import (
	"net/http"
	"strings"

	"github.com/AmyangXYZ/sgo"
)

// CORSOpt is options of CORS middleware.
type CORSOpt struct {
	Skipper      Skipper
	AllowOrigins []string
	AllowMethods []string
	AllowHeaders []string
}

// CORS returns a Cross-Origin Resource Sharing (CORS) middleware.
func CORS(opt CORSOpt) sgo.HandlerFunc {
	return func(ctx *sgo.Context) error {
		if opt.Skipper == nil {
			opt.Skipper = DefaultSkipper
		}
		if opt.AllowOrigins == nil {
			opt.AllowOrigins = []string{"*"}
		}
		if opt.AllowMethods == nil {
			opt.AllowMethods = []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions}
		}
		if opt.AllowHeaders == nil {
			opt.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
		}

		if opt.Skipper(ctx) {
			ctx.Next()
			return nil
		}
		allowOrigins := strings.Join(opt.AllowOrigins, ",")
		allowMethods := strings.Join(opt.AllowMethods, ",")
		allowHeaders := strings.Join(opt.AllowHeaders, ",")

		ctx.Resp.Header().Set("Access-Control-Allow-Origin", allowOrigins)
		ctx.Resp.Header().Set("Access-Control-Allow-Methods", allowMethods)
		ctx.Resp.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		ctx.Next()
		return nil
	}
}
