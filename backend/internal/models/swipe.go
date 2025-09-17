package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Swipe records a user's swipe action on another user's profile.
type Swipe struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Swiper    primitive.ObjectID `bson:"swiper" json:"swiper" validate:"required"`
	Swiped    primitive.ObjectID `bson:"swiped" json:"swiped" validate:"required"`
	Direction string             `bson:"direction" json:"direction" validate:"required,oneof=right left"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}
