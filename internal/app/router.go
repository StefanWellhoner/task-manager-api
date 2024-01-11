package app

import (
	"github.com/StefanWellhoner/task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *services.GormDatabase) *gin.Engine {
	router := gin.Default()

	return router
}
