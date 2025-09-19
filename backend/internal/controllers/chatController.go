package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"sportmatch-backend/internal/models"
	"sportmatch-backend/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for now, but in production, restrict to your frontend domain
		return true
	},
}

// ServeWs handles websocket connections for chat.
func ServeWs(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection to websocket: %v", err)
		return
	}
	defer conn.Close()

	userID, exists := c.Get("userID")
	if !exists {
		log.Println("Unauthorized websocket connection: User ID not found in context")
		conn.WriteJSON(gin.H{"error": "Unauthorized"})
		return
	}
	currentUserID := userID.(string)

	log.Printf("User %s connected via websocket", currentUserID)

	// TODO: Manage active websocket connections (e.g., store in a map)
	// For now, just read messages and echo them back (or save to DB)
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message from %s: %v", currentUserID, err)
			break
		}

		// Assuming p contains a JSON message with sender, receiver, and message content
		var chatInput models.ChatMessageInput
		if err := bson.UnmarshalExtJSON(p, true, &chatInput); err != nil {
			log.Printf("Error unmarshalling chat message: %v", err)
			continue
		}

		receiverID, err := primitive.ObjectIDFromHex(chatInput.ReceiverID)
		if err != nil {
			log.Printf("Invalid receiver ID in chat message: %v", err)
			continue
		}

		senderObjID, err := primitive.ObjectIDFromHex(currentUserID)
		if err != nil {
			log.Printf("Invalid sender ID from context: %v", err)
			continue
		}

		newChat := models.Chat{
			ID:        primitive.NewObjectID(),
			Sender:    senderObjID,
			Receiver:  receiverID,
			Message:   chatInput.Message,
			Timestamp: time.Now(),
			Read:      false, // Mark as unread initially
		}

		_, err = database.MI.Db.Collection("chats").InsertOne(context.TODO(), newChat)
		if err != nil {
			log.Printf("Failed to save chat message to DB: %v", err)
			continue
		}

		log.Printf("Message from %s to %s: %s", currentUserID, chatInput.ReceiverID, chatInput.Message)

		// TODO: Send message to the receiver's active websocket connection if available
		// For now, just echo back to sender
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Error writing message back to %s: %v", currentUserID, err)
			break
		}
	}
}

// GetUserConversations retrieves a list of conversations for the authenticated user.
func GetUserConversations(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Find distinct users that the current user has chatted with
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"$or": []bson.M{
					{"sender": objID},
					{"receiver": objID},
				},
			},
		},
		{
			"$sort": bson.M{"timestamp": -1}, // Sort by latest message
		},
		{
			"$group": bson.M{
				"_id": "$$ROOT.sender", // Group by sender to find distinct conversations
				"lastMessage": bson.M{"$first": "$$ROOT.message"},
				"lastMessageTime": bson.M{"$first": "$$ROOT.timestamp"},
				"otherUser": bson.M{"$first": bson.M{
					"$cond": []interface{}{
						bson.M{"$eq": []interface{}{"$$ROOT.sender", objID}},
						"$$ROOT.receiver",
						"$$ROOT.sender",
					},
				}},
			},
		},
		{
			"$lookup": bson.M{
				"from":         "users",
				"localField":   "otherUser",
				"foreignField": "_id",
				"as":           "otherUserDetails",
			},
		},
		{
			"$unwind": "$otherUserDetails",
		},
		{
			"$project": bson.M{
				"_id":             0,
				"id":              "$otherUser",
				"userName":        "$otherUserDetails.name",
				"userAvatar":      bson.M{"$arrayElemAt": []interface{}{"$otherUserDetails.profilePictureUrls", 0}}, // Get first profile picture
				"lastMessage":     1,
				"lastMessageTime": 1,
				// TODO: Calculate unreadCount
				"unreadCount": 0,
			},
		},
	}

	cursor, err := database.MI.Db.Collection("chats").Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Printf("Error aggregating conversations: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve conversations"})
		return
	}
	defer cursor.Close(context.TODO())

	var conversations []gin.H
	if err = cursor.All(context.TODO(), &conversations); err != nil {
		log.Printf("Error decoding conversations: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode conversations"})
		return
	}

	c.JSON(http.StatusOK, conversations)
}

// GetChatMessages retrieves messages for a specific conversation.
func GetChatMessages(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	objUserID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	otherUserID := c.Param("otherUserID")
	objOtherUserID, err := primitive.ObjectIDFromHex(otherUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid other user ID format"})
		return
	}

	// Find messages between current user and other user
	filter := bson.M{
		"$or": []bson.M{
			{"sender": objUserID, "receiver": objOtherUserID},
			{"sender": objOtherUserID, "receiver": objUserID},
		},
	}

	cursor, err := database.MI.Db.Collection("chats").Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Error finding chat messages: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}
	defer cursor.Close(context.TODO())

	var messages []models.Chat
	if err = cursor.All(context.TODO(), &messages); err != nil {
		log.Printf("Error decoding chat messages: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode messages"})
		return
	}

	// Optionally mark messages as read
	_, err = database.MI.Db.Collection("chats").UpdateMany(
		context.TODO(),
		bson.M{"sender": objOtherUserID, "receiver": objUserID, "read": false},
		bson.M{"$set": bson.M{"read": true, "updatedAt": time.Now()}},
	)
	if err != nil {
		log.Printf("Error marking messages as read: %v", err)
	}

	c.JSON(http.StatusOK, messages)
}
