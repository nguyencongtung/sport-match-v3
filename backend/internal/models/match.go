package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Match (Event) represents a scheduled sporting activity that users can create and join.
type Match struct {
	ID                primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Creator           primitive.ObjectID   `bson:"creator" json:"creator"` // Reference to the user who created the match
	Sport             string               `bson:"sport" json:"sport"`
	StartTime         time.Time            `bson:"startTime" json:"startTime"`
	MaxPeople         int                  `bson:"maxPeople" json:"maxPeople"`
	CurrentParticipants int                `bson:"currentParticipants" json:"currentParticipants"`
	Fee               float64              `bson:"fee" json:"fee"`
	Location          MatchLocation        `bson:"location" json:"location"`
	Level             string               `bson:"level" json:"level"`
	Description       string               `bson:"description,omitempty" json:"description,omitempty"`
	Participants      []primitive.ObjectID `bson:"participants" json:"participants"` // List of users who have joined the match
	CreatedAt         time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt         time.Time            `bson:"updatedAt" json:"updatedAt"`
}

// MatchLocation defines details about the match location.
type MatchLocation struct {
	Name        string    `bson:"name" json:"name"`
	Address     string    `bson:"address,omitempty" json:"address,omitempty"`
	Coordinates []float64 `bson:"coordinates,omitempty" json:"coordinates,omitempty"` // [longitude, latitude]
}

// CreateMatchInput defines the input structure for creating a new match.
type CreateMatchInput struct {
	Sport       string        `json:"sport" binding:"required"`
	StartTime   time.Time     `json:"startTime" binding:"required"`
	MaxPeople   int           `json:"maxPeople" binding:"required,min=1"`
	Fee         float64       `json:"fee"`
	Location    MatchLocation `json:"location" binding:"required"`
	Level       string        `json:"level" binding:"required"`
	Description string        `json:"description,omitempty"`
}

// JoinMatchInput defines the input structure for joining a match.
type JoinMatchInput struct {
	MatchID string `json:"matchId" binding:"required"`
}
