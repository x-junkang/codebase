package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		path := c.FullPath()
		// before request

		c.Next()

		// after request
		latency := time.Since(t)

		// access the status we are sending
		status := c.Writer.Status()
		log.Printf("path: %s; latency: %dms; code: %d", path, latency.Milliseconds(), status)
	}
}
