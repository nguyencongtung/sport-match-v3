package main

import (
	"context"
	"log"
	"os"

	"sportmatch-backend/internal/config"
	"sportmatch-backend/pkg/database"

	"sportmatch-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.ConnectDB()
	defer func() {
		if err := database.MI.Client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()

	// Define API routes here
	routes.SetupUserRoutes(router, database.MI.Db)
	// swipeRoutes.SetupSwipeRoutes(router, database.MI.Db)
	// chatRoutes.SetupChatRoutes(router, database.MI.Db)
	// matchRoutes.SetupMatchRoutes(router, database.MI.Db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
