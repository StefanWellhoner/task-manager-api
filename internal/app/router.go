package app

import (
	"fmt"

	"github.com/StefanWellhoner/task-manager-api/internal/config"
	"github.com/StefanWellhoner/task-manager-api/internal/handlers"
	"github.com/StefanWellhoner/task-manager-api/internal/middleware"
	"github.com/StefanWellhoner/task-manager-api/internal/services"
	"github.com/gin-gonic/gin"
)

// setupMiddleware sets up the middleware for the router.
func setupMiddleware(router *gin.Engine) {
	router.Use(middleware.RequestID())
	router.Use(middleware.Logger())
}

func registerRoutes(router *gin.Engine) {
	db, err := services.New(config.Get())

	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %v", err))
	}

	authGroup := router.Group("/auth")
	{
		authGroup.GET("/refresh", handlers.Refresh)
		authGroup.POST("/login", handlers.Login)
		authGroup.POST("/logout", handlers.Logout)
		authGroup.POST("/register", handlers.Register(db))
	}

	userGroup := router.Group("/users")
	{
		userGroup.GET("/", handlers.GetUsers)
		userGroup.GET("/:id", handlers.GetUser)
		userGroup.PUT("/:id", handlers.UpdateUser)
		userGroup.DELETE("/:id", handlers.DeleteUser)
	}

	taskGroup := router.Group("/tasks")
	{
		taskGroup.GET("/", handlers.GetTasks)
		taskGroup.POST("/", handlers.CreateTask)
		taskGroup.GET("/:id", handlers.GetTask)
		taskGroup.PUT("/:id", handlers.UpdateTask)
		taskGroup.DELETE("/:id", handlers.DeleteTask)

		taskGroup.POST("/:id/assign", handlers.AssignTask)
		taskGroup.POST("/:id/unassign", handlers.UnassignTask)
		taskGroup.POST("/:id/reassign", handlers.ReassignTask)

		taskGroup.GET("/:id/deadlines", handlers.GetTaskDeadlines)
		taskGroup.POST("/:id/reminders", handlers.CreateTaskReminder)

		taskGroup.GET("/filter", handlers.FilterTasks)
		taskGroup.GET("/search", handlers.SearchTasks)
	}

	categoryGroup := router.Group("/categories")
	{
		categoryGroup.GET("/", handlers.GetCategories)
		categoryGroup.POST("/", handlers.CreateCategory)
		categoryGroup.GET("/:id", handlers.GetCategory)
		categoryGroup.PUT("/:id", handlers.UpdateCategory)
		categoryGroup.DELETE("/:id", handlers.DeleteCategory)
		categoryGroup.GET("/:id/tasks", handlers.GetCategoryTasks)
	}
}

// SetupRouter sets up the router with all the necessary routes and middleware.
func SetupRouter() *gin.Engine {
	router := gin.New()

	router.SetTrustedProxies(nil)

	setupMiddleware(router)
	registerRoutes(router)

	return router
}
