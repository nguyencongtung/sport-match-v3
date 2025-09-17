package routes

import (
	"sportmatch-backend/internal/controllers"
	"sportmatch-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupUserRoutes sets up user-related API routes
func SetupUserRoutes(router *gin.Engine, db *mongo.Database) {
	userController := controllers.NewUserController(db.Collection("users"))

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userController.RegisterUser)
		userRoutes.POST("/login", userController.LoginUser)
		userRoutes.GET("/profile", middleware.AuthMiddleware(), userController.GetUserProfile)
		userRoutes.PUT("/profile", middleware.AuthMiddleware(), userController.UpdateUserProfile)
	}
}
