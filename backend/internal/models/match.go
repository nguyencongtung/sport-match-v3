package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Match (Event) represents a scheduled sporting activity that users can create and join.
type Match struct {
	ID                  primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Creator             primitive.ObjectID   `bson:"creator" json:"creator" validate:"required"`
	Sport               string               `bson:"sport" json:"sport" validate:"required"`
	StartTime           time.Time            `bson:"startTime" json:"startTime" validate:"required"`
	MaxPeople           int                  `bson:"maxPeople" json:"maxPeople" validate:"required,min=1"`
	CurrentParticipants int                  `bson:"currentParticipants" json:"currentParticipants"`
	Fee                 float64              `bson:"fee" json:"fee"`
	Location            MatchLocation        `bson:"location" json:"location" validate:"required"`
	Level               string               `bson:"level" json:"level" validate:"required,oneof=Beginner Intermediate Advanced All Levels"`
	Description         string               `bson:"description,omitempty" json:"description,omitempty"`
	Participants        []primitive.ObjectID `bson:"participants" json:"participants"`
	CreatedAt           time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt           time.Time            `bson:"updatedAt" json:"updatedAt"`
}

// MatchLocation represents the location details for a match.
type MatchLocation struct {
	Name        string    `bson:"name" json:"name" validate:"required"`
	Address     string    `bson:"address,omitempty" json:"address,omitempty"`
	Coordinates []float64 `bson:"coordinates,omitempty" json:"coordinates,omitempty"` // [longitude, latitude]
}
