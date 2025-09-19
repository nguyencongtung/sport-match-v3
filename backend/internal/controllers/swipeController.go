package controllers

import (
	"context"
	"net/http"
	"time"

	"sportmatch-backend/internal/models"
	"sportmatch-backend/pkg/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateSwipe handles a user's swipe action.
func CreateSwipe(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	swiperID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid swiper ID format"})
		return
	}

	var input models.SwipeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	swipedID, err := primitive.ObjectIDFromHex(input.SwipedID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid swiped user ID format"})
		return
	}

	// Prevent self-swiping
	if swiperID == swipedID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot swipe on yourself"})
		return
	}

	newSwipe := models.Swipe{
		ID:        primitive.NewObjectID(),
		Swiper:    swiperID,
		Swiped:    swipedID,
		Direction: input.Direction,
		Timestamp: time.Now(),
	}

	_, err = database.MI.Db.Collection("swipes").InsertOne(context.TODO(), newSwipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record swipe"})
		return
	}

	// Check for mutual match if direction is "right"
	if input.Direction == "right" {
		var mutualSwipe models.Swipe
		err := database.MI.Db.Collection("swipes").FindOne(context.TODO(), bson.M{
			"swiper":    swipedID,
			"swiped":    swiperID,
			"direction": "right",
		}).Decode(&mutualSwipe)

		if err == nil {
			// Mutual like found! This is a connection.
			// TODO: Implement connection logic (e.g., create a chat room, send notification)
			c.JSON(http.StatusOK, gin.H{"message": "Swipe recorded, it's a connection!", "is_match": true})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Swipe recorded", "is_match": false})
}
