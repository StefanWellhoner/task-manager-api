package handlers

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func HandleResponse(c *gin.Context, statusCode int, message string, details any) {
	c.JSON(statusCode, Response{Code: statusCode, Message: message, Details: details})
}
