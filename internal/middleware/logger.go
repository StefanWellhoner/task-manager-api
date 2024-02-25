package middleware

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		logEntry := map[string]interface{}{
			"method:":      c.Request.Method,
			"path":         c.Request.URL.Path,
			"address":      c.ClientIP(),
			"responseTime": time.Since(start).String(),
			"userAgent":    c.Request.Header.Get("User-Agent"),
			"status":       c.Writer.Status(),
			"content-type": c.Writer.Header().Get("Content-Type"),
			"requestID":    c.GetString(RequestIDHeader),
		}

		log.Println(toJson(logEntry))
	}
}

func toJson(data map[string]interface{}) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Logger middleware failed to marshal log entry: %v", err)
		return "{}"
	}
	return string(bytes)
}
