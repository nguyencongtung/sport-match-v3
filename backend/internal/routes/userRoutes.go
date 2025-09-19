package routes

import (
	"sportmatch-backend/internal/controllers"
	"sportmatch-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes sets up user-related API routes
func SetupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.POST("/login", controllers.LoginUser)
		userRoutes.GET("/profile", middleware.AuthMiddleware(), controllers.GetUserProfile)
		userRoutes.PUT("/profile", middleware.AuthMiddleware(), controllers.UpdateUserProfile)
		userRoutes.POST("/profile/pictures", middleware.AuthMiddleware(), controllers.UploadProfilePictures)
	}
}
