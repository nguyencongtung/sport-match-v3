package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Chat represents a single message within a conversation between two connected users.
type Chat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Sender    primitive.ObjectID `bson:"sender" json:"sender" validate:"required"`
	Receiver  primitive.ObjectID `bson:"receiver" json:"receiver" validate:"required"`
	Message   string             `bson:"message" json:"message" validate:"required"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	Read      bool               `bson:"read" json:"read"`
}
