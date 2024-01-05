package app

import (
	"net/http"

	"github.com/StefanWellhoner/task-manager-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		// Index routes
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Index"})
		})

		// Task routes
		tasks := v1.Group("/tasks")
		{
			tasks.GET("", handlers.GetTasks)
			tasks.POST("", handlers.CreateTask)
			tasks.DELETE("", handlers.DeleteTask)
			tasks.PUT("", handlers.UpdateTask)
		}

		// User routes
		users := v1.Group("/users")
		{
			users.POST("/register", handlers.RegisterUser)
			users.POST("/login", handlers.LoginUser)
		}
	}

	return router
}
