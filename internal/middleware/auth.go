package middleware

import (
	"fmt"
	"strings"

	"github.com/StefanWellhoner/task-manager-api/internal/config"
	"github.com/StefanWellhoner/task-manager-api/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(config.Get().Secrets.Jwt)

func extractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no authorization header provided")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return bearerToken[1], nil
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token expired")
			}
		}

		return nil, err
	}

	return token, nil
}

func parseTokenClaims(token *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token or claims")
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractToken(c)
		if err != nil {
			handlers.HandleResponse(c, 401, "Unauthorized", err.Error())
			c.Abort()
			return
		}

		token, err := validateToken(tokenString)
		if err != nil {
			handlers.HandleResponse(c, 401, "Unauthorized", err.Error())
			c.Abort()
			return
		}

		claims, err := parseTokenClaims(token)
		if err != nil {
			handlers.HandleResponse(c, 401, "Unauthorized", err.Error())
			c.Abort()
			return
		}

		if claims["token_type"] != "access" {
			handlers.HandleResponse(c, 401, "Unauthorized", "invalid token type")
			c.Abort()
			return
		}

		c.Set("userID", claims["user_id"])
		c.Next()
	}
}
