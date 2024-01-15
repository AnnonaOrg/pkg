package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		//check for incomint header,use it if exists
		requestId := c.Request.Header.Get("X-Request-Id")

		//Create request id with UUID4
		if requestId == "" {
			requestId = uuid.New().String()
		}

		//Expose it for use in the application
		c.Set("X-Request-Id", requestId)

		//Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
