package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents an individual user of the SportMatch application.
type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string             `bson:"name" json:"name" validate:"required"`
	Email             string             `bson:"email" json:"email" validate:"required,email,unique"`
	PasswordHash      string             `bson:"passwordHash" json:"passwordHash" validate:"required,min=8"`
	DateOfBirth       time.Time          `bson:"dateOfBirth" json:"dateOfBirth" validate:"required"`
	Gender            string             `bson:"gender" json:"gender" validate:"required,oneof=Male Female Other/Prefer not to say"`
	ProfilePictureUrls []string           `bson:"profilePictureUrls,omitempty" json:"profilePictureUrls,omitempty"`
	Sports            []string           `bson:"sports" json:"sports" validate:"required,min=1"`
	LookingForGender  string             `bson:"lookingForGender" json:"lookingForGender" validate:"required,oneof=Men Women Everyone/Any"`
	AgeRange          struct {
		Min int `bson:"min" json:"min" validate:"required,min=18"`
		Max int `bson:"max" json:"max" validate:"required,min=18"`
	} `bson:"ageRange" json:"ageRange" validate:"required"`
	SkillLevel        string             `bson:"skillLevel" json:"skillLevel" validate:"required,oneof=Beginner Intermediate Advanced All Levels"`
	CreatedAt         time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt         time.Time          `bson:"updatedAt" json:"updatedAt"`
}
