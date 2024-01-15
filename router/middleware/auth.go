package middleware

import (
	"github.com/AnnonaOrg/pkg/errno"
	"github.com/AnnonaOrg/pkg/handler"
	"github.com/AnnonaOrg/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
