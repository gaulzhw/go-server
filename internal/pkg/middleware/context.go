package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/gaulzhw/go-server/internal/pkg/constant"
	"github.com/gaulzhw/go-server/internal/pkg/log"
)

// UsernameKey defines the key in gin context which represents the owner of the secret.

// Context is a middleware that injects common prefix fields to gin.Context.
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(log.KeyRequestID, c.GetString(constant.XRequestIDKey))
		c.Set(log.KeyUsername, c.GetString(constant.XUsernameKey))
		c.Next()
	}
}
