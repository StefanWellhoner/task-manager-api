package middleware

import "github.com/gin-gonic/gin"

func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Content Security Policy
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; object-src 'none'; frame-ancestors 'none'; form-action 'self'; base-uri 'self';")

		// X-Frame-Options
		c.Writer.Header().Set("X-Frame-Options", "DENY")

		// X-Content-Type-Options
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		// X-XSS-Protection
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")

		// Referrer-Policy
		c.Writer.Header().Set("Referrer-Policy", "no-referrer-when-downgrade")

		// Feature-Policy (Consider using Permissions-Policy depending on the most recent browser support)
		c.Writer.Header().Set("Feature-Policy", "geolocation 'none'; microphone 'none'; camera 'none'")

		c.Next()
	}
}
