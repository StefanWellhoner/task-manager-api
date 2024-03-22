package handlers

import (
	"net/http"

	"github.com/StefanWellhoner/task-manager-api/internal/errors"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if serviceError, ok := err.(*errors.ServiceError); ok {
		status := serviceError.Status
		if status == 0 {
			status = http.StatusInternalServerError
		}

		c.JSON(status, gin.H{
			"type":    serviceError.Type,
			"message": serviceError.Message,
			"status":  status,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"type":    errors.InternalError,
		"message": "An internal error occurred",
		"status":  http.StatusInternalServerError,
	})
}
