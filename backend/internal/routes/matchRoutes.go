package routes

import (
	"sportmatch-backend/internal/controllers"
	"sportmatch-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupMatchRoutes sets up match-related API routes
func SetupMatchRoutes(router *gin.Engine) {
	matchRoutes := router.Group("/matches")
	matchRoutes.Use(middleware.AuthMiddleware()) // All match routes require authentication
	{
		matchRoutes.POST("/", controllers.CreateMatch)
		matchRoutes.GET("/", controllers.GetMatches)
		matchRoutes.POST("/:matchID/join", controllers.JoinMatch)
	}
}
