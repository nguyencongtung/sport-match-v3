package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Swipe represents a user's swipe action on another user's profile.
type Swipe struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Swiper    primitive.ObjectID `bson:"swiper" json:"swiper"`       // Reference to the user who performed the swipe
	Swiped    primitive.ObjectID `bson:"swiped" json:"swiped"`       // Reference to the user whose profile was swiped on
	Direction string             `bson:"direction" json:"direction"` // "right" for like, "left" for pass
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}

// SwipeInput defines the input structure for a new swipe.
type SwipeInput struct {
	SwipedID  string `json:"swipedId" binding:"required"`
	Direction string `json:"direction" binding:"required,oneof=right left"`
}
