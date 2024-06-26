package cors

import (
	"github.com/gin-gonic/gin"
	"github.com/ginx-contribs/ginx/constant/methods"
	"github.com/ginx-contribs/ginx/constant/status"
	"github.com/rs/cors"
)

// Options is a configuration container to setup the CORS middleware.
type Options = cors.Options

// corsWrapper is a wrapper of cors.Cors handler which preserves information
// about configured 'optionPassthrough' option.
type corsWrapper struct {
	*cors.Cors
	optionPassthrough bool
}

// build transforms wrapped cors.Cors handler into Gin middleware.
func (c corsWrapper) build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
		if !c.optionPassthrough &&
			ctx.Request.Method == methods.Options &&
			ctx.GetHeader("Access-Control-Request-Method") != "" {
			// Abort processing next Gin middlewares.
			ctx.AbortWithStatus(status.OK.Code())
		}
	}
}

// AllowAll creates a new CORS Gin middleware with permissive configuration
// allowing all origins with all standard methods with any header and
// credentials.
func AllowAll() gin.HandlerFunc {
	return corsWrapper{Cors: cors.AllowAll()}.build()
}

// Default creates a new CORS Gin middleware with default options.
func Default() gin.HandlerFunc {
	return corsWrapper{Cors: cors.Default()}.build()
}

// New creates a new CORS Gin middleware with the provided options.
func New(options Options) gin.HandlerFunc {
	return corsWrapper{cors.New(options), options.OptionsPassthrough}.build()
}
