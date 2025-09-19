package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"sportmatch-backend/internal/models"
	"sportmatch-backend/pkg/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateMatch handles the creation of a new sporting event/match.
func CreateMatch(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	creatorID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid creator ID format"})
		return
	}

	var input models.CreateMatchInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.StartTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Start time must be in the future"})
		return
	}

	newMatch := models.Match{
		ID:                  primitive.NewObjectID(),
		Creator:             creatorID,
		Sport:               input.Sport,
		StartTime:           input.StartTime,
		MaxPeople:           input.MaxPeople,
		CurrentParticipants: 1, // Creator is the first participant
		Fee:                 input.Fee,
		Location:            input.Location,
		Level:               input.Level,
		Description:         input.Description,
		Participants:        []primitive.ObjectID{creatorID},
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	_, err = database.MI.Db.Collection("matches").InsertOne(context.TODO(), newMatch)
	if err != nil {
		log.Printf("Error creating match: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create match"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Match created successfully", "match_id": newMatch.ID.Hex()})
}

// GetMatches retrieves a list of upcoming sporting events/matches.
func GetMatches(c *gin.Context) {
	// Optional filters from query parameters
	sportFilter := c.Query("sport")
	levelFilter := c.Query("level")

	filter := bson.M{
		"startTime": bson.M{"$gt": time.Now()}, // Only upcoming matches
	}

	if sportFilter != "" {
		filter["sport"] = sportFilter
	}
	if levelFilter != "" {
		filter["level"] = levelFilter
	}

	findOptions := options.Find().SetSort(bson.D{{Key: "startTime", Value: 1}}) // Sort by start time ascending

	cursor, err := database.MI.Db.Collection("matches").Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Printf("Error finding matches: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve matches"})
		return
	}
	defer cursor.Close(context.TODO())

	var matches []models.Match
	if err = cursor.All(context.TODO(), &matches); err != nil {
		log.Printf("Error decoding matches: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode matches"})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// JoinMatch allows an authenticated user to join an existing match.
func JoinMatch(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	participantID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid participant ID format"})
		return
	}

	matchID := c.Param("matchID")
	objMatchID, err := primitive.ObjectIDFromHex(matchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match ID format"})
		return
	}

	var match models.Match
	err = database.MI.Db.Collection("matches").FindOne(context.TODO(), bson.M{"_id": objMatchID}).Decode(&match)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match not found"})
		return
	}

	if match.CurrentParticipants >= match.MaxPeople {
		c.JSON(http.StatusConflict, gin.H{"error": "Match is already full"})
		return
	}

	// Check if user is already a participant
	for _, pID := range match.Participants {
		if pID == participantID {
			c.JSON(http.StatusConflict, gin.H{"error": "User is already a participant in this match"})
			return
		}
	}

	update := bson.M{
		"$inc": bson.M{"currentParticipants": 1},
		"$push": bson.M{"participants": participantID},
		"$set":  bson.M{"updatedAt": time.Now()},
	}

	_, err = database.MI.Db.Collection("matches").UpdateByID(context.TODO(), objMatchID, update)
	if err != nil {
		log.Printf("Error joining match: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to join match"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully joined match"})
}
