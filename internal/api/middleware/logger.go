package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("[REQUEST] %s %s | %v", c.Request.Method, c.Request.RequestURI, duration)
	}
}
