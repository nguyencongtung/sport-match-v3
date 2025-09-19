package routes

import (
	"sportmatch-backend/internal/controllers"
	"sportmatch-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupSwipeRoutes sets up swipe-related API routes
func SetupSwipeRoutes(router *gin.Engine) {
	swipeRoutes := router.Group("/swipes")
	swipeRoutes.Use(middleware.AuthMiddleware()) // All swipe routes require authentication
	{
		swipeRoutes.POST("/", controllers.CreateSwipe)
		// TODO: Add routes for getting potential matches (users to swipe on)
	}
}
