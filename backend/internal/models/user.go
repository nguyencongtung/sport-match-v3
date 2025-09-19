package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents an individual user of the SportMatch application.
type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string             `bson:"name" json:"name"`
	Email             string             `bson:"email" json:"email"`
	PasswordHash      string             `bson:"passwordHash" json:"-"` // - to omit from JSON output
	DateOfBirth       time.Time          `bson:"dateOfBirth" json:"dateOfBirth"`
	Gender            string             `bson:"gender" json:"gender"`
	ProfilePictureUrls []string           `bson:"profilePictureUrls,omitempty" json:"profilePictureUrls,omitempty"`
	Sports            []string           `bson:"sports" json:"sports"`
	LookingForGender  string             `bson:"lookingForGender" json:"lookingForGender"`
	AgeRange          AgeRange           `bson:"ageRange" json:"ageRange"`
	SkillLevel        string             `bson:"skillLevel" json:"skillLevel"`
	CreatedAt         time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt         time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// AgeRange defines the preferred age range for connections/matches.
type AgeRange struct {
	Min int `bson:"min" json:"min"`
	Max int `bson:"max" json:"max"`
}

// UserRegisterInput defines the input structure for user registration.
type UserRegisterInput struct {
	Name            string    `json:"name" binding:"required"`
	Email           string    `json:"email" binding:"required,email"`
	Password        string    `json:"password" binding:"required,min=6"`
	ConfirmPassword string    `json:"confirmPassword" binding:"required"`
	DateOfBirth     time.Time `json:"dateOfBirth" binding:"required"`
	Gender          string    `json:"gender" binding:"required"`
}

// UserLoginInput defines the input structure for user login.
type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserProfileUpdateInput defines the input structure for updating user profile dynamic attributes.
type UserProfileUpdateInput struct {
	Sports           []string `json:"sports"`
	LookingForGender string   `json:"lookingForGender"`
	AgeRange         AgeRange `json:"ageRange"`
	SkillLevel       string   `json:"skillLevel"`
}
