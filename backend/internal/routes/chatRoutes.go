package routes

import (
	"sportmatch-backend/internal/controllers"
	"sportmatch-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupChatRoutes sets up chat-related API routes
func SetupChatRoutes(router *gin.Engine) {
	chatRoutes := router.Group("/chats")
	chatRoutes.Use(middleware.AuthMiddleware()) // All chat routes require authentication
	{
		chatRoutes.GET("/", controllers.GetUserConversations)
		chatRoutes.GET("/:otherUserID/messages", controllers.GetChatMessages)
		chatRoutes.GET("/ws", controllers.ServeWs) // WebSocket endpoint
	}
}
