package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"sportmatch-backend/internal/models"
	"sportmatch-backend/internal/utils"
	"sportmatch-backend/pkg/database"
	"sportmatch-backend/pkg/jwt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RegisterUser handles new user registration.
func RegisterUser(c *gin.Context) {
	var input models.UserRegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email already exists
	var existingUser models.User
	err := database.MI.Db.Collection("users").FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create new user
	newUser := models.User{
		ID:           primitive.NewObjectID(),
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: hashedPassword,
		DateOfBirth:  input.DateOfBirth,
		Gender:       input.Gender,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		// Default values for dynamic attributes, to be updated later
		Sports:            []string{},
		LookingForGender:  "Everyone",
		AgeRange:          models.AgeRange{Min: 18, Max: 99},
		SkillLevel:        "All Levels",
		ProfilePictureUrls: []string{},
	}

	_, err = database.MI.Db.Collection("users").InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(newUser.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "token": token, "user_id": newUser.ID.Hex()})
}

// LoginUser handles user login.
func LoginUser(c *gin.Context) {
	var input models.UserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user models.User
	err := database.MI.Db.Collection("users").FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "user_id": user.ID.Hex()})
}

// GetUserProfile retrieves the profile of the authenticated user.
func GetUserProfile(c *gin.Context) {
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

	var user models.User
	err = database.MI.Db.Collection("users").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Do not return password hash
	user.PasswordHash = ""
	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile handles updating dynamic user profile attributes.
func UpdateUserProfile(c *gin.Context) {
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

	var input models.UserProfileUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"sports":           input.Sports,
			"lookingForGender": input.LookingForGender,
			"ageRange":         input.AgeRange,
			"skillLevel":       input.SkillLevel,
			"updatedAt":        time.Now(),
		},
	}

	_, err = database.MI.Db.Collection("users").UpdateByID(context.TODO(), objID, update)
	if err != nil {
		log.Printf("Error updating user profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// UploadProfilePictures handles uploading profile pictures.
func UploadProfilePictures(c *gin.Context) {
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

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	files := form.File["photos"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No photos uploaded"})
		return
	}

	var imageUrls []string
	for _, file := range files {
		// In a real application, you would upload this file to a cloud storage
		// like AWS S3, Google Cloud Storage, or a similar service.
		// For this MVP, we'll simulate the upload and store a placeholder URL.
		// The actual URL would be returned by the cloud storage service.

		// Example: Save file locally (for demonstration purposes only)
		// dst := "uploads/" + file.Filename
		// if err := c.SaveUploadedFile(file, dst); err != nil {
		// 	log.Printf("Failed to save file locally: %v", err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		// 	return
		// }
		// imageUrls = append(imageUrls, "/"+dst) // Store local path

		// Placeholder for cloud storage URL
		placeholderURL := "https://example.com/profile_pictures/" + file.Filename
		imageUrls = append(imageUrls, placeholderURL)
	}

	update := bson.M{
		"$set": bson.M{
			"profilePictureUrls": imageUrls,
			"updatedAt":          time.Now(),
		},
	}

	_, err = database.MI.Db.Collection("users").UpdateByID(context.TODO(), objID, update)
	if err != nil {
		log.Printf("Error updating profile picture URLs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile pictures"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile pictures uploaded successfully", "urls": imageUrls})
}
