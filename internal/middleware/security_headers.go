package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Content Security Policy
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; object-src 'none'; frame-ancestors 'none'; form-action 'self'; base-uri 'self';")

		// Cross-Origin Resource Policy
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// X-Frame-Options
		c.Writer.Header().Set("X-Frame-Options", "DENY")

		// X-Content-Type-Options
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		// X-XSS-Protection
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")

		// Referrer-Policy
		c.Writer.Header().Set("Referrer-Policy", "no-referrer-when-downgrade")

		c.Next()
	}
}
