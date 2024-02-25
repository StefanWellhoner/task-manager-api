package handlers

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func HandleError(c *gin.Context, statusCode int, message string, details any) {
	c.JSON(statusCode, ErrorResponse{Code: statusCode, Message: message, Details: details})
}
