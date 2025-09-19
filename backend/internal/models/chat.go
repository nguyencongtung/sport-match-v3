package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Chat represents a single message within a conversation between two connected users.
type Chat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Sender    primitive.ObjectID `bson:"sender" json:"sender"`     // Reference to the user who sent the message
	Receiver  primitive.ObjectID `bson:"receiver" json:"receiver"` // Reference to the user who received the message
	Message   string             `bson:"message" json:"message"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	Read      bool               `bson:"read" json:"read"` // Indicates if the message has been read by the receiver
}

// ChatMessageInput defines the input structure for sending a new chat message.
type ChatMessageInput struct {
	ReceiverID string `json:"receiverId" binding:"required"`
	Message    string `json:"message" binding:"required"`
}
