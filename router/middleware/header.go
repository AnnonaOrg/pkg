package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//Nocache is a middleware function that appends headers
//to pervent the client from caching the Http response
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-status,max-age=0,must-revalidate,value")
	c.Header("Expires", "The ,01 jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

//Options is a middleware function that appends headers
//for options requests and aborts then exits the middleware
//chain and ends the request.
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		withCredentials := "true"
		origin := c.Request.Header.Get("Origin")
		if len(origin) == 0 {
			origin = "*"
			withCredentials = "false"
		}
		c.Header("Access-Control-Allow-Credentials", withCredentials)
		c.Header("Access-Control-Allow-Origin", origin)
		// c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept, x-action-addr")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

//secure is a middleware function that appends security
//and resource access headers.
func Secure(c *gin.Context) {
	withCredentials := "true"
	origin := c.Request.Header.Get("Origin")
	if len(origin) == 0 {
		origin = "*"
		withCredentials = "false"
	}
	// fmt.Println(origin)
	// for k, v := range c.Request.Header {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(c.Request.URL)
	c.Header("Access-Control-Allow-Credentials", withCredentials)
	c.Header("Access-Control-Allow-Origin", origin)
	// c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/X-XSS-Protection
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
	// Also consider adding Content-Security-Policy headers
	// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
}
